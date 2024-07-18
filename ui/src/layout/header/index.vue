<template>
  <el-row :gutter="20">
    <el-col :span="6"><div class="grid-content ep-bg-purple" />1</el-col>
    <el-col :span="6"><div class="grid-content ep-bg-purple" />2</el-col>
    <el-col :span="6"><div class="grid-content ep-bg-purple" />3</el-col>
    <el-col :span="6"><div class="grid-content ep-bg-purple" />4</el-col>
  </el-row>
</template>

<style scoped lang="scss">
.el-row {
  margin-bottom: 20px;
}
.el-row:last-child {
  margin-bottom: 0;
}
.el-col {
  border-radius: 4px;
}

.grid-content {
  border-radius: 4px;
  min-height: 36px;
}
</style>