<script setup lang="tsx">
import { NButton, NPopconfirm, NTag } from 'naive-ui';
import { fetchGetUserList, deleteUser } from '@/service/api';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { enableStatusRecord, userGenderRecord } from '@/constants/business';
import { useTable, useTableOperate } from '@/hooks/common/table';
import UserOperateDrawer from './modules/user-operate-drawer.vue';
import UserSearch from './modules/user-search.vue';
import { ref } from "vue";
import AssignRoleModel from "./modules/assign-role-model.vue";
import ResetPasswdModel from "./modules/reset-passwd-model.vue";
import { useBoolean } from "~/packages/hooks";
import { useAuth } from '@/hooks/business/auth';

const { hasAuth } = useAuth();
const appStore = useAppStore();
const {
  columns,
  columnChecks,
  data,
  getData,
  getDataByPage,
  loading,
  mobilePagination,
  searchParams,
  resetSearchParams
} = useTable({
  apiFn: fetchGetUserList,
  showTotal: true,
  apiParams: {
    current: 1,
    size: 10,
    // if you want to use the searchParams in Form, you need to define the following properties, and the value is null
    // the value can not be undefined, otherwise the property in Form will not be reactive
    status: null,
    username: null,
    gender: null,
    nickname: null,
    phone: null,
    email: null
  },
  columns: () => [
    {
      type: 'selection',
      align: 'center',
      width: 48
    },
    {
      key: 'index',
      title: $t('common.index'),
      align: 'center',
      width: 64
    },
    {
      key: 'username',
      title: $t('page.manage.user.userName'),
      align: 'center',
      minWidth: 100
    },
    {
      key: 'gender',
      title: $t('page.manage.user.userGender'),
      align: 'center',
      width: 100,
      render: row => {
        if (row.gender === null) {
          return null;
        }

        const tagMap: Record<Api.SystemManage.UserGender, NaiveUI.ThemeColor> = {
          1: 'primary',
          2: 'error'
        };

        const label = $t(userGenderRecord[row.gender]);

        return <NTag type={tagMap[row.gender]}>{label}</NTag>;
      }
    },
    {
      key: 'nickname',
      title: $t('page.manage.user.nickName'),
      align: 'center',
      minWidth: 100
    },
    {
      key: 'phone',
      title: $t('page.manage.user.userPhone'),
      align: 'center',
      width: 120
    },
    {
      key: 'email',
      title: $t('page.manage.user.userEmail'),
      align: 'center',
      minWidth: 200
    },
    {
      key: 'status',
      title: $t('page.manage.user.userStatus'),
      align: 'center',
      width: 100,
      render: row => {
        if (row.status === null) {
          return null;
        }

        const tagMap: Record<Api.Common.EnableStatus, NaiveUI.ThemeColor> = {
          1: 'success',
          2: 'warning'
        };

        const label = $t(enableStatusRecord[row.status]);

        return <NTag type={tagMap[row.status]}>{label}</NTag>;
      }
    },
    {
      key: 'operate',
      title: $t('common.operate'),
      align: 'center',
      width: 350,
      render: row => (
        <div class="flex-center gap-8px">
          {hasAuth('sys:user:edit') && (
            <NButton type="primary" ghost size="small" onClick={() => edit(row.id)}>
              {$t('common.edit')}
            </NButton>
          )}
          {hasAuth('sys:user:assignRoles') && (
            <NButton type="primary" ghost size="small" onClick={() => handleRoleAuth(row.id, row.roles)}>
              {$t('page.manage.user.roleAuth')}
            </NButton>
          )}
          {hasAuth('sys:user:resetPwd') && (
            <NButton type="primary" ghost size="small" onClick={() => handleResetPwd(row.id)}>
              {$t('page.manage.user.resetPassWd')}
            </NButton>
          )}
          {hasAuth('sys:user:delete') && (
            <NPopconfirm onPositiveClick={() => handleDelete(row.id)}>
              {{
                default: () => $t('common.confirmDelete'),
                trigger: () => (
                  <NButton type="error" ghost size="small">
                    {$t('common.delete')}
                  </NButton>
                )
              }}
            </NPopconfirm>
          )}
        </div>
      )
    }
  ]
});

const {
  drawerVisible,
  operateType,
  editingData,
  handleAdd,
  handleEdit,
  checkedRowKeys,
  onBatchDeleted,
  onDeleted
  // closeDrawer
} = useTableOperate(data, getData);

async function handleBatchDelete() {
  // request
  console.log(checkedRowKeys.value);

  onBatchDeleted();
}

async function handleDelete(id: number) {
  // request
  const { error } = await deleteUser(id);
  if (error) {
    return;
  }

  await onDeleted();
}

function edit(id: number) {
  handleEdit(id);
}

const roleAuthUserId = ref<number>(0);
const roleList = ref<Api.SystemManage.Role[]>([]);
const { bool: roleAuthVisible, setTrue: openRoleAuthModal } = useBoolean();
function handleRoleAuth(id: number, roles: Api.SystemManage.Role[]) {
  roleAuthUserId.value = id;
  roleList.value = roles;
  openRoleAuthModal();
}

const { bool: resetPwdVisible, setTrue: openResetPwdModal } = useBoolean();
function handleResetPwd(id: number) {
  roleAuthUserId.value = id;
  openResetPwdModal();
}

</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <UserSearch v-model:model="searchParams" @reset="resetSearchParams" @search="getDataByPage" />
    <NCard :title="$t('page.manage.user.title')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <template #header-extra>
        <TableHeaderOperation
          v-model:columns="columnChecks"
          :display-add-btn="hasAuth('sys:user:add')"
          :disabled-delete="checkedRowKeys.length === 0"
          :loading="loading"
          @add="handleAdd"
          @delete="handleBatchDelete"
          @refresh="getData"
        />
      </template>
      <NDataTable
        v-model:checked-row-keys="checkedRowKeys"
        :columns="columns"
        :data="data"
        size="small"
        :flex-height="!appStore.isMobile"
        :scroll-x="962"
        :loading="loading"
        remote
        :row-key="row => row.id"
        :pagination="mobilePagination"
        class="sm:h-full"
      />
      <UserOperateDrawer
        v-model:visible="drawerVisible"
        :operate-type="operateType"
        :row-data="editingData"
        @submitted="getDataByPage"
      />
      <AssignRoleModel v-model:visible="roleAuthVisible" :user-id="roleAuthUserId" :roles="roleList" @submitted="getDataByPage"/>
      <ResetPasswdModel v-model:visible="resetPwdVisible" :user-id="roleAuthUserId" />
    </NCard>
  </div>
</template>

<style scoped></style>
