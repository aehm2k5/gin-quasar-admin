<template>
    <q-card>
        <div ref="linechart" style="height: 400px;"></div>
        <q-resize-observer @resize="onResize" />
    </q-card>
</template>

<script setup>
import { markRaw, onMounted, ref } from 'vue';
const echarts = require('echarts')
const chart = ref(null)
const linechart = ref(null)

onMounted(() => {
    init()
})
const init = () => {
    let ct = linechart.value;
    echarts.dispose(ct);
    chart.value = markRaw(echarts.init(ct));
    chart.value.setOption(options.value);
}
const onResize = () => {
    chart.value.resize()
}

const options = ref({
    xAxis: {
        type: 'category',
        data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
    },
    yAxis: {
        type: 'value'
    },
    series: [
        {
            data: [820, 932, 901, 934, 1290, 1330, 1320],
            type: 'line',
            smooth: true
        }
    ]
})

</script>
