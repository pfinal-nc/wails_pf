//Global JS function for greeting
function greet() {
    //Get user input
    let inputName = document.getElementById("name").value;

    //Call Go Greet function
    window.go.main.App.Greet(inputName).then(result => {
        //Display result from Go
        document.getElementById("result").innerHTML = result;
    }).catch(err => {
        console.log(err);
    }).finally(() => {
        console.log("finished!")
    });
}

function event_cpu_on() {
    let cpuChart = echarts.init(document.getElementById("cpu"), 'dark', {
        renderer: 'canvas',
        useDirtyRect: false,
    });
    let cpu_option = {
        series: [
            {
                type: 'gauge',
                axisLine: {
                    lineStyle: {
                        width: 10,
                        color: [
                            [0.3, '#67e0e3'],
                            [0.7, '#37a2da'],
                            [1, '#fd666d']
                        ]
                    }
                },
                pointer: {
                    itemStyle: {
                        color: 'auto'
                    }
                },
                axisTick: {
                    distance: -10,
                    length: 8,
                    lineStyle: {
                        color: '#fff',
                        width: 2
                    }
                },
                splitLine: {
                    distance: -10,
                    length: 15,
                    lineStyle: {
                        color: '#fff',
                        width: 3
                    }
                },
                axisLabel: {
                    color: 'inherit',
                    distance: 10,
                    fontSize: 10
                },
                detail: {
                    valueAnimation: true,
                    formatter: '{value}%',
                    color: 'inherit'
                },
                data: [
                    {
                        value: 70
                    }
                ]
            }
        ]
    }
    runtime.EventsOn("cpu_usage", function (cpu_usage) {
        console.log(cpu_usage.avg)
        cpuChart.setOption({
            series: [
                {
                    data: [
                        {
                            name:"CPU使用率",
                            value: cpu_usage.avg.toFixed(2)
                        }
                    ]
                }
            ]
        })
    })
    cpu_option && cpuChart.setOption(cpu_option);
    window.addEventListener('resize', function () {
        cpuChart.resize()
    });
}
function event_mem_on() {
    let mem_used = 0;
    runtime.EventsOn("mem_usage", function (mem_usage) {
        console.log(mem_usage.used_percent)

    })
    layui.use(function(){
        let element = layui.element;
        let util = layui.util;

    })
}
$(function () {
    event_cpu_on()
    event_mem_on()
})