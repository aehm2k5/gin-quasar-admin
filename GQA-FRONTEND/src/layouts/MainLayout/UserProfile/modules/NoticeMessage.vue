<template>
    <div>
        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">
            <template v-slot:body-cell-notice_type="props">
                <q-td :props="props">
                    <GqaDictShow dictName="noticeType" :dictCode="props.row.notice_type" />
                </q-td>
            </template>

            <template v-slot:body-cell-notice_read="props">
                <q-td :props="props">
                    <GqaDictShow dictName="statusYesNo"
                        :dictCode="props.row.notice_to_user.filter(item => item.to_user === username)[0].user_read" />
                </q-td>
            </template>

            <template v-slot:body-cell-notice_sent="props">
                <q-td :props="props">
                    <GqaDictShow dictName="statusYesNo" :dictCode="props.row.notice_sent" />
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="primary" @click="readNotice(props.row)" :label="$t('Read')" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <NoticeDetail ref="noticeDetail" @hide="hide" />
    </div>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { DictOptions } from 'src/utils/dict'
import { FormatDateTime } from 'src/utils/date'
import { useUserStore } from 'src/stores/user'
import NoticeDetail from './NoticeDetail.vue'

const $q = useQuasar()
const { t } = useI18n()
const userStore = useUserStore()
const url = {
    list: 'notice/get-notice-list',
}
const columns = computed(() => {
    return [
        { name: 'id', align: 'center', label: 'ID', field: 'id' },
        { name: 'notice_title', align: 'center', label: t('Title'), field: 'notice_title' },
        { name: 'notice_type', align: 'center', label: t('Type'), field: 'notice_type' },
        { name: 'notice_read', align: 'center', label: t('Read') + t('Status'), field: 'notice_read' },
        { name: 'notice_sent', align: 'center', label: t('Sent'), field: 'notice_sent' },
        { name: 'actions', align: 'center', label: t('Actions'), field: 'actions' },
    ]
})
const {
    pagination,
    queryParams,
    pageOptions,
    GqaDictShow,
    GqaAvatar,
    loading,
    tableData,
    recordDetailDialog,
    showAddForm,
    showEditForm,
    onRequest,
    handleSearch,
    resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

const username = computed(() => userStore.GetUsername())

onMounted(() => {
    queryParams.value = {
        notice_type: 'message',
        notice_sent: 'yes',
        notice_to_user: String(username.value),
    }
    pagination.value.sortBy = 'created_at'
    onRequest({
        pagination: pagination.value,
        queryParams: queryParams.value
    })
})
const noticeDetail = ref(null)
const readNotice = (row) => {
    noticeDetail.value.show(row)
}
const hide = () => {
    onRequest({
        pagination: pagination.value,
        queryParams: queryParams.value
    })
}
</script>
