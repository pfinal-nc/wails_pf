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

function event_host_on() {
    window.go.main.App.HostInfo().then(result => {
        //Display result from Go
        res = JSON.parse(result)
        console.log(res)
        document.getElementById("hostname").textContent = res.hostname
        document.getElementById("core").textContent = res.platform
    }).catch(err => {
        console.log(err);
    }).finally(() => {
        console.log("finished!")
    });
}

function event_cpu_on() {
    layui.use(function () {
        runtime.EventsOn("cpu_usage", function (cpu_usage) {
            // element.progress('demo-filter-progress', cpu_usage.avg + '%'); // 设置 50% 的进度
            document.getElementById("used").textContent = cpu_usage.avg + '% '
        })
    })

    window.go.main.App.CpuInfo().then(result => {
        //Display result from Go
        res = JSON.parse(result)
        document.getElementById("cpu_num").textContent = res.cpu_number
    }).catch(err => {
        console.log(err);
    }).finally(() => {
        console.log("finished!")
    });
}

function event_mem_on() {
    let mem_used = 0;
    runtime.EventsOn("mem_usage", function (mem_usage) {
        //console.log(mem_usage.used_percent.toFixed(2))
    })
    window.go.main.App.MemInfo().then(result => {
        //Display result from Go
        res = JSON.parse(result)
        console.log(res)
        //document.getElementById("cpu_num").textContent = res.cpu_number
    }).catch(err => {
        console.log(err);
    }).finally(() => {
        console.log("finished!")
    });
}

function event_time() {
    // 获取页面id=time的元素（P元素<p id="time">12:12:12 AM</p>）
    let time = document.querySelector("#time");
    let day = document.querySelector("#date")

    // 定义函数传入time元素
    function up(time_el, dateDom) {
        // 获取当前时间
        let date = new Date();
        // 获取时分秒
        let h = date.getHours();
        let m = date.getMinutes();
        let s = date.getSeconds();
        // 上午与下午
        let day_night = "AM";

        // 计算上午与下午
        if (h > 12) {
            h = h - 12;
            day_night = "PM";
        }

        // 如果时间小于10则前面补充0
        if (h < 10) {
            h = "0" + h;
        }
        if (m < 10) {
            m = "0" + m;
        }
        if (s < 10) {
            s = "0" + s;
        }

        // 拼接时间并且赋值给time元素的文本中，从而显示
        time_el.textContent = h + ":" + m + ":" + s + " " + day_night;

        let year = date.getFullYear()
        let month = date.getMonth() + 1
        let day = date.getDate()
        let week = '';
        switch (date.getDay()) {
            case 1:
                week = "星期一"
                break;
            case 2:
                week = "星期二"
                break;
            case 3:
                week = "星期三"
                break;
            case 4:
                week = "星期四"
                break;
            case 5:
                week = "星期五"
                break;
            case 6:
                week = "星期六"
                break;
            case 0:
                week = "星期日"
                break;
        }
        dateDom.textContent = year + "年" + month + "月" + day + "日  " + week

    }

    // 定时器，每秒执行一次实现更新
    setInterval(() => {
        up(time, day);
    }, 1000);
}

$(function () {
    event_host_on()
    event_cpu_on()
    event_time()
    live2d_settings['modelId'] = 1; live2d_settings['modelTexturesId'] = 87;
    initModel("/libs/live2d/assets/waifu-tips.json")

})