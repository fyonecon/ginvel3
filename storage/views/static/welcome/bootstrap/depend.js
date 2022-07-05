/*主入口js，下面一般不要做修改。**/
/*https://www.sojson.com/aaencode.html*/
/*2019/9/12/9:55*/

// 原生依赖
const index_load = {
    "index_js": [
        "routes/route.js",
        "bootstrap/md5.js",
        "bootstrap/clipboard.js",
        "static/js/swiper-bundle.min.js",
        "static/js/jquery.fullPage.js",
        // 以上文件不需要更改位置
        "config/must.js", // 根据路由加载模块页面，并启动全局安全校验

    ],
};

// 浏览器环境检查，主要检测是否支持ES6语法
(function () {
    try{
        let check = 1;
        new Promise(function(resolve, reject) {
            if (check === 1){
                check++;
                resolve();
            } else {
                check--;
                reject();
            }
        }).then(function () {
            const map = new Map([
                ["check", check],
            ]);
            map.get("check");
        });
    }catch (e) {
        console.error(e);
        alert("该古老的浏览器不支持ES6语法，需更换浏览器。");
        window.location.replace(index_file_url + "parts/help/help-es6.html");
    }
})();

// 检测白名单refer
(function () {
    let the_host = window.location.host;
    let the_refer = document.referrer;

    try {
        let check_refer = app_refer[0].check_refer;
        let jump_site = app_refer[0].jump_site;
        let white_refer = app_refer[0].white_refer;
        // view.log([the_host, the_refer]);

        if (!check_refer){
            view.log("refer检测已关闭");
        }else {
            if (the_refer){
                the_refer = the_refer.slice(0, 30);
                view.log("refer检测已开启，len=30");
                for (let j=0; j<white_refer.length; j++){
                    let the_white_refer = white_refer[j];
                    if (view.string_include_string(the_refer, the_white_refer) !== -1){ // 处在白名单
                        //view.log("白名单refer=" + the_white_refer);
                    }else {
                        if (jump_site){
                            window.location.replace(jump_site);
                        }
                        break;
                    }
                }
            }else {
                view.log("refer为空时跳过refer校验");
            }
        }
    }catch (e) {
        view.log([the_host, the_refer, '参数不全，已跳过']);
    }

})();

// 监听url是否发生变化，确保页面跳转成功
(function () {
    let url = window.location.href;

    window.onhashchange = function () {
        let now_url = window.location.href;

        // 只获取和对比#号之前的网址
        let index1 = url.lastIndexOf("\#");
        if (index1>0){
            url = url.substring(0, index1);
        }

        let index2 = now_url.lastIndexOf("\#");
        if (index2>0){
            now_url = now_url.substring(0, index2);
        }

        if (url && url !== now_url){
            // console.log(["url不同：", url, now_url]);
            window.location.reload();
        }else {
            console.log(["跳过url变化检测（只对比#号之前的）：", url, now_url]);
        }

    };
})();


// 加载框架模块文件
(function () {

    const depend = {  // 依赖函数
        "get_url_param": function (url, key) {  // 获取url中的参数
            let url_str = "";
            if(!url){
                url_str = window.location.href;
            }else {
                url_str = url;
            }
            // 正则匹配url中的参数，如果存在键，则返回键的值，如果不存在则返回null
            let regExp = new RegExp("([?]|&|#)" + key + "=([^&|^#]*)(&|$|#)");
            let result = url_str.match(regExp);
            if (result) {
                return decodeURIComponent(result[2]); // 转义还原参数
            } else {
                return "";
            }
        },
        "page_file": function (pages_index) {  // 添加页面js、css资源文件
            if (pages_index === null){
                console.error("route.js数组参数找不到，页面和框架数据将不能渲染。");
                setTimeout(function () {
                    window.location.replace(route_404); // 则进入默认页
                }, 1000);
                return;
            }

            let head = document.head || document.getElementsByTagName("head")[0];
            let file = pages[pages_index].file[0];

            // css不需要异步
            for (let i=0; i<file.css.length; i++){
                let link = document.createElement('link');
                link.setAttribute("href", file_url + file.css[i] +"?"+ page_time);
                link.setAttribute("rel", "stylesheet");
                head.appendChild(link);
            }

            let js_all = [];
            for (let i=0; i<file.js.length; i++){
                let the_p = new Promise((resolve, reject) => {
                    let script = document.createElement("script");
                    script.setAttribute("src", file_url + file.js[i] +"?"+ page_time);
                    head.appendChild(script);
                    script.onload = function () {resolve(i); };
                });
                js_all.push(the_p);
            }
            Promise.all(js_all).then((result) => {
                depend.page_all_js_has();
            }).catch((error) => {
                console.error(error);
            });
        },
        "page_all_js_has": function () {  // 页面全部js加载完后执行
            view.log("Files Cache_time = "+cache_time +"s");

            document.getElementById("loading-div").classList.add("hide");
            time_loaded = Math.floor((new Date()).getTime());
            let view_loaded_time = time_loaded - time_start;

            try {
                let head = document.head || document.getElementsByTagName("head")[0];
                let script = document.createElement("script");
                script.setAttribute("src", file_url + "static/js/page_loaded.js?"+page_time);
                head.appendChild(script);
                script.onload = function () {
                    view.log('page_loaded');
                };
            }catch (e) {
                console.error(e);
                console.log("=error=page_loaded=");
            }

            let page = depend.get_url_param("", "route");
            try {
                must_safe_check([
                    view_loaded_time,
                    "框架解析完成，用时"+view_loaded_time+"ms", "开始执行"+page+"页面数据>>",
                    index_file_url,
                    index_depend_file,
                ]);
            }catch (e) {
                console.error("错误提示：情况1：【可忽略】must_safe_check()" + "页面起始模块函数未定义，但是此函数可忽略。情况2：must_safe_check()函数缺失，请参考如下报错：");
                console.error(e);
            }
        },
        'load_page': function () { // 加载页面
            let page_name = "";     // 拉取哪个html文件块
            let _file = "";         // 真实文件路径+文件名
            let pages_index = null; // 页面资源索引

            // 匹配路由名
            let p1 = new Promise((resolve, reject) => {
                page_name = depend.get_url_param("", "route");
                for (let i=0; i<pages.length; i++){ // 获取真正文件路径名
                    if (pages[i].route === page_name){
                        _file = page_url + "pages/" + pages[i].file_path + "?"+page_time;
                        document.getElementsByTagName("title")[0].innerHTML = pages[i].title;
                        pages_index = i;
                        resolve('找到值');
                    }else if (pages.length-1 === i) {
                        resolve('未找到值');
                    }
                }
            });

            // 处理路由页面名
            let p2 = new Promise((resolve, reject) => {
                if (page_name === ""){ // 空路由
                    window.location.replace(route_default);  // 则进入默认页
                    resolve('默认页');
                }else {
                    if (pages_index === null){ // 未匹配路由
                        console.error("页面没有正确路由?route=xxx，将进入默认页面。");
                        time_error = Math.floor((new Date()).getTime());
                        setTimeout(function () {
                            window.location.replace(route_404);  // 则进入404页
                        },0);
                        resolve('未匹配');
                    }else{
                        resolve('已匹配');
                    }
                }
            });

            // 获取页面txt
            let p3 = new Promise((resolve, reject) =>{
                $.ajax({ // 利用ajax的get请求获取文本内容
                    url: _file,
                    async: true,
                    success: function (data) {
                        let div = document.createElement("div");
                        div.classList.add("app-page");
                        div.classList.add("page-div");
                        div.classList.add("clear");
                        div.setAttribute("id", "app-page");
                        div.setAttribute("data-view", ""+view.js_rand(1000000000000, 99999999999999));
                        div.classList.add("page-div-" + view.js_rand(1000000000000, 99999999999999));
                        div.innerHTML = data;

                        let depend = document.getElementById("depend");
                        depend.classList.add("depend-div-" + view.js_rand(1000000000000, 99999999999999));
                        depend.setAttribute("data-view", ""+view.js_rand(1000000000000, 99999999999999));

                        depend.appendChild(div); // 将模块渲染入主文件

                        resolve('文本已请求');
                    },
                    error: function (error) {
                        console.error("缺失模块html文件=" + error);
                        console.error("1.非同源政策限制模块文件的拉取；2.本应用需要服务器环境（网络环境）；3.htm组件文件404。");
                        time_error = Math.floor((new Date()).getTime());
                        setTimeout(function () {
                            window.location.replace(index_file_url + "parts/help/help-html.html");
                        },1000);

                        resolve('缺失模块html文件');
                    }
                });
            });

            // 渲染最后页面的资源
            Promise.all([p1, p2, p3]).then((result) => {
                let head = document.head || document.getElementsByTagName("head")[0];

                // 页面渲染完毕，开始执行公共css、js引入
                for (let i=0; i<page_static_file.css.length; i++){
                    let link = document.createElement('link');
                    link.setAttribute("href", file_url + page_static_file.css[i] + "?" + page_time);
                    link.setAttribute("rel", "stylesheet");
                    head.appendChild(link);
                }

                let js_all = [];
                for (let i=0; i<page_static_file.js.length; i++){
                    let the_p = new Promise((resolve, reject) => {
                        let script = document.createElement("script");
                        script.setAttribute("src", file_url + page_static_file.js[i] + "?" + page_time);
                        head.appendChild(script);
                        script.onload = function () {resolve(i); };
                    });
                    js_all.push(the_p);
                }

                Promise.all(js_all).then((result) => {
                    depend.page_file(pages_index);
                }).catch((error) => {
                    console.error(error);
                });

            }).catch((error) => {
                console.error(error);
            });

        },

    };

    // 校验文件引入参数是否已经存在，不存在就不需要解析框架
    if( typeof time_start === "undefined" || typeof file_url === "undefined" || typeof page_url === "undefined" || typeof page_time === "undefined" ){
        console.error("参数未定义：%s，框架产生了异步时差，需要决解框架Bug。5s秒后将重试网页。", [time_start, file_url, page_url, page_time]);
        setTimeout(function () {
            window.location.reload();
        }, 5000);
    }else{
        // 引入原生依赖
        let head = document.head || document.getElementsByTagName("head")[0];
        let js_all = [];
        for (let i=0; i<index_load.index_js.length; i++){
            let the_p = new Promise((resolve, reject) => {
                let pages_script = document.createElement("script");
                pages_script.setAttribute("src", file_url + index_load.index_js[i]+"?" + page_time);
                head.appendChild(pages_script);
                pages_script.onload = function () {resolve(i); };
            });
            js_all.push(the_p);
        }

        Promise.all(js_all).then((result) => {
            depend.load_page();
        }).catch((error) => {
            console.error(error);
        });

    }

})();



