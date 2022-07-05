/*
 * 局部模块js
 * */


const search_debug = false; // 调试日志，false关闭日志，true显示日志
const title = " 👈 简洁主页"; // 当前页面标题
const search = [ // 搜索引擎列表，分为移动和PC、前缀和后缀。自定义。
    {
        "name": "Bing国际",
        "m-url": "?route=search&engine=bing&word=",
        "pc-url": "?route=search&engine=bing&word=",
        "url_right": "",
    },
    {
        "name": "Baidu中文",
        "m-url": "?route=search&engine=baidu&word=",
        "pc-url": "?route=search&engine=baidu&word=",
        "url_right": "",
    },
    {
        "name": "Google", // 引擎名称，可视5个字
        "m-url": "?route=search&engine=google&word=", // 移动端，前缀
        "pc-url": "?route=search&engine=google&word=", // PC端，前缀
        "url_right": "", // 参数的固顶后缀
    },
    {
        "name": "搜英文电影",
        "m-url": "?route=search&engine=video&word=",
        "pc-url": "?route=search&engine=video&word=",
        "url_right": "",
    },

    {
        "name": "搜收费音乐",
        "m-url": "?route=search&engine=music&word=",
        "pc-url": "?route=search&engine=music&word=",
        "url_right": "",
    },

    {
        "name": "英中翻译",
        "m-url": "https://fanyi.baidu.com/translate#en/zh/",
        "pc-url": "https://fanyi.baidu.com/translate#en/zh/",
        "url_right": "",
    },
    {
        "name": "中英翻译",
        "m-url": "https://fanyi.baidu.com/translate#zh/en/",
        "pc-url": "https://fanyi.baidu.com/translate#zh/en/",
        "url_right": "",
    },
    {
        "name": "搜Github",
        "m-url": "https://github.com/search?&type=Repositories&q=",
        "pc-url": "https://github.com/search?&type=Repositories&q=",
        "url_right": "",
    },
    {
        "name": "搜电影MP4",
        "m-url": "https://m.domp4.com/",
        "pc-url": "https://www.domp4.com/",
        "url_right": "blank",
    },
    {
        "name": "查询IPv4",
        "m-url": "http://ip.tool.chinaz.com/",
        "pc-url": "http://ip.tool.chinaz.com/",
        "url_right": "",
    },
    {
        "name": "查询SEO",
        "m-url": "https://seo.chinaz.com/",
        "pc-url": "https://seo.chinaz.com/",
        "url_right": "",
    },
    {
        "name": "查询Whois",
        "m-url": "http://whois.chinaz.com/",
        "pc-url": "http://whois.chinaz.com/",
        "url_right": "",
    },

    {
        "name": "CSDN博客",
        "m-url": "https://so.csdn.net/so/search/s.do?q=",
        "pc-url": "https://so.csdn.net/so/search/s.do?q=",
        "url_right": "",
    },
    {
        "name": "cnblogs博客",
        "m-url": "https://zzk.cnblogs.com/s/blogpost?w=",
        "pc-url": "https://zzk.cnblogs.com/s/blogpost?w=",
        "url_right": "",
    },

];



// 调试日志
function console_log(txt) {
    search_debug === true ? console.info(txt) : "";
}

/*
 * 实现自定义的N次连续点击
 * many_click(_click_num, call_func)
 * 必填：_click_num 点击次数 [1, 10]
 * 必填：call_func 回调函数
 * 选填：_id 是长按手势传入的目标标签id
 * */
let click_before_time = 0;
let click_num = 0;

function many_click(_click_num, call_func, _id) {
    if (!call_func) { console.info("many_click(_click_num, call_func)无回调函数"); return; }
    if (_click_num === "long") { /*实现长按*/
        if (!_id) { console.info("_id为必填。many_click('long', call_func, _id)"); return; }
        long_press(_id, call_func);
        return;
    }
    // 安全校验
    if (typeof _click_num !== "number") { console.info("many_click(_click_num, call_func)的点击次数为number类型"); return; }
    // 处理click_num的新值情况
    if (click_num === 0) {
        click_num = _click_num;
    } else {
        if (click_num < 1 || click_num > 10) { click_num = 1; } /*只准1击至10击，其他情况默认1击*/
    }
    // 处理点击之时差
    let click_time = Date.parse(new Date()) + (new Date()).getMilliseconds(); // 毫秒时间戳
    if ((click_time - click_before_time) < 400) { // 下一次点击是否成功
        click_before_time = Date.parse(new Date()) + (new Date()).getMilliseconds();
        click_num--;
    } else { // 第一次点击
        click_before_time = Date.parse(new Date()) + (new Date()).getMilliseconds();
        if (click_num < _click_num) { /*清除历史不成功点击的参数*/
            click_num = _click_num;
        }
    }
    // N次成功点击后启用回调函数，并初始化click_num
    if (click_num === 1) {
        call_func("回调函数不需要传参");
        click_num = 0; /*初始化点击次数*/
    }
}

/*
 * 长按事件
 * long_press(_id, call_func)
 * */
function long_press(_id, call_func) {
    let timer = null;
    _id.addEventListener("touchstart", function() {
        timer = setTimeout(function() {
            call_func(_id);
        }, 1200); // 定义长按时间
    });
    _id.addEventListener("touchend", function() {
        clearTimeout(timer);
    });
}

//写入cookies
// time = 1*24*60*60*1000;
function setCookie(name, value, time) {
    if (!time) {
        time = 100 * 24 * 60 * 60 * 1000; // 默认100天
    }
    var exp = new Date();
    exp.setTime(exp.getTime() + time);
    document.cookie = name + "=" + escape(value) + ";expires=" + exp.toGMTString();
}
//读取cookies
function getCookie(name) {
    var arr, reg = new RegExp("(^| )" + name + "=([^;]*)(;|$)");
    if (arr = document.cookie.match(reg)) {
        return unescape(arr[2]);
    } else {
        return null;
    }
}
//删除cookies
function delCookie(name) {
    var exp = new Date();
    exp.setTime(exp.getTime() - 1);
    var cval = getCookie(name);
    if (cval != null) {
        document.cookie = name + "=" + cval + ";expires=" + exp.toGMTString();
    }
}

// 二维码图位置<img class="qr_img" id="qr-img" />
function make_new_qr(content, width, height, call_func, id) {
    let _content = content ? content : "没有设置二维码参数";
    let _width = width ? width : 120;
    let _height = height ? height : 120;
    try {
        document.getElementById("qrcode").remove(); // 每次都移除老的
    } catch (e) {
        console_log("不存在qrcode-div节点");
    }
    let div = document.createElement("div");
    div.classList.add("qrcode");
    div.style.display = "none";
    div.setAttribute("id", "qrcode");
    document.getElementById("content").appendChild(div);
    let qrcode = new QRCode("qrcode", {
        text: _content,
        width: _width,
        height: _height,
        colorDark: "#000000",
        colorLight: '#f5f5f5',
        correctLevel: QRCode.CorrectLevel.L
    });

    setTimeout(function() {
        try {
            call_func(1, id);
        } catch (e) {
            console_log([call_func, id]);
        }
    }, 100);

}

function show_qr(num, id) {
    let el = document.getElementById(id);
    let img = document.getElementsByClassName("qr_img")[0].getAttribute("src");

    console_log([num, el]);

    if (!img) {
        try {
            document.getElementsByClassName("div-qr")[0].classList.add("hide");
        } catch (e) {
            console_log("img空=1=");
        }
    } else {
        console_log(img);

        // el.setAttribute("src", img);
        el.style.backgroundImage = "url('" + img + "')";

        let _width = window.innerWidth;
        let _screen_width = screen.width;
        resize_width(_width, _screen_width);

    }

}

function get_url_param(url, key) { // 获取url中的参数
    // 兼容模式url地址，例如：poop.html?page=3&ok=222#p=2#name=kd
    let url_str = "";
    if (!url) { url_str = window.location.href; } else { url_str = url; }
    let regExp = new RegExp("([?]|&|#)" + key + "=([^&|^#]*)(&|$|#)");
    let result = url_str.match(regExp);
    if (result) {
        return decodeURIComponent(result[2]); // 转义还原参数
    } else {
        return ""; // 没有匹配的键即返回空
    }
}

function timestamp() {
    return new Date().getTime();
}



const search_cookie_pre = "search_";
const search_eq = search_cookie_pre + "_eq";
let search_time_style = 0; // 自动校正iframe
let focus_time = 3 * 60 * 60 * 1000; // 保护用户输入框隐私，3h不聚焦删一次
let blur_time = 6 * 60 * 60 * 1000; // 保护用户输入框隐私，6h聚焦删一次
let dead_input_num = 0; // 自动初始化输入框

function set_search(val) { // 配置当前的搜索引擎
    console_log("配置当前搜索引擎");
    setCookie(search_eq, val, 30 * 24 * 60 * 60 * 1000);
    for (let i = 0; i < document.getElementsByClassName("option").length; i++) {
        document.getElementsByClassName("option")[i].removeAttribute("selected");
    }
    document.getElementsByClassName("option-" + val)[0].setAttribute("selected", "selected");
    document.getElementsByTagName("title")[0].innerText = document.getElementsByClassName("option-" + val)[0].innerText + title;
}

function create_input(pre) { // 渲染模板
    console_log("渲染模板数据");

    document.getElementsByTagName("title")[0].innerText = title;
    let content = document.getElementsByClassName("content")[0];
    content.innerHTML = '<div class="input-div" id="input-div"><select class="select search-style select-none" id="select"></select><input type="text" value="" maxlength="500" id="input" class="input search-style"  placeholder="' + pre + '输入内容，双击Enter搜索"/><div class="clear"></div></div><div class="input-history-div" id="input-history"></div><div class="clear"></div><div class="search-btn-div" id="search-btn"></div><div class="res-div"></div>';
    let append_tag = [];
    for (let i = 0; i < search.length; i++) {
        let tag = '<option class="option option-' + i + '" value="' + i + '">' + search[i]["name"] + '</option>';
        append_tag.push(tag);
    }
    document.getElementsByClassName("select")[0].innerHTML = append_tag.join("");

    document.getElementById("input-div").classList.add("input-div-blur");

    let _eq = getCookie(search_eq);
    if (_eq) { set_search(_eq); } else { set_search(0); }

    setTimeout(function() {
        delete_loading();
        write_tips_text('若浏览器阻止打开新标签，务必手动选择允许打开');
        // make_new_qr(window.location.href, 200, 200, show_qr, "img-show_qr");
    }, 300);

}

function dead_input(del_time) { // 处理清空用户输入的情况
    try {
        clearTimeout(dead_input_num);
        console_log(dead_input_num + "-清除成功");
    } catch (e) {
        console_log(dead_input_num + "-timeout is none");
    }
    dead_input_num = setTimeout(function() {
        create_input("重新");
        console_log(del_time);
    }, del_time);
    console_log(dead_input_num);
}

function run_search() { // 执行搜索
    update_history();
    change_blur();
    try {
        clearInterval(search_time_style);
    } catch (e) {
        console_log("第一次进入页面是没有定时器的");
    }
    let _select = document.getElementById("select");
    let engine = _select.options[_select.selectedIndex].value;
    let _input = document.getElementById("input").value;
    if (!_input.trim()) {
        console_log("内容不能为空");
        // view.alert_txt("搜索内容不能为空", 1500);
        view.notice_txt("搜索内容不能为空", 1500);
        change_focus();
        return;
    }

    let reg = /^([hH][tT]{2}[pP]:\/\/|[hH][tT]{2}[pP][sS]:\/\/)+([A-Za-z0-9-~\/])/; // 至少是 http://a 这种格式
    if (!reg.test(_input)) {
        console_log("不是网址");
        _input = encodeURIComponent(_input);
    } else {
        console_log("是网址");
        window.open(_input, "_blank"); // 搜索4/4
        return;
    }

    let url_right = search[engine]["url_right"].trim(); // 参数固定后缀
    let m_url = "";
    let pc_url = "";
    if (url_right === "blank") { // 对于有些网站，只能打开主页
        m_url = search[engine]["m-url"]; // get，移动端
        pc_url = search[engine]["pc-url"]; // get，PC端
    } else { // 正常搜索
        m_url = search[engine]["m-url"] + _input + url_right; // get，移动端
        pc_url = search[engine]["pc-url"] + _input + url_right; // get，PC端
    }
    let tab_url = "";

    if (window.innerWidth > 640) {
        write_tips_text("PC模式会自动打开新标签来展示搜索结果");
        tab_url = pc_url;
    } else {
        // 操作iOS设备Bug情况
        let u = navigator.userAgent;
        let isAndroid = u.indexOf('Android') > -1 || u.indexOf('Adr') > -1;
        let isiOS = !!u.match(/\(i[^;]+;( U;)? CPU.+Mac OS X/);
        if (isAndroid === true || isiOS === false) { // android
            tab_url = m_url;
            console_log("Android");
        } else if (isAndroid === false || isiOS === true) { // ios
            console_log("iOS");
            write_tips_text("iOS移动设备会自动打开新标签来展示搜索结果");
            tab_url = m_url;
        } else { // pc
            tab_url = pc_url;
            console_log("PC");
        }
    }

    show_loading();
    write_tips_text('已经在新标签打开了本次搜索结果');
    change_blur(); // 主动退去键盘

    setTimeout(function() {
        show_history();
        console_log("打开新标签也买你");
        window.open(tab_url, "_blank");
    }, 500);
    setTimeout(function() {
        delete_loading();
    }, 1200);

}

function init_dom() {

    // 初始化页面输入框
    create_input("");

    // 初始化搜索按钮
    document.getElementById("search-btn").innerHTML = '' +
        //
        '<div class="search-btn-center do-btn-center must-btn select-none">' +
        '   <span class="search-btn-style history-btn-span click red" title="清空搜索历史记录">清空历史</span>' +
        '   <span class="search-btn-style refresh-btn-span click green" title="清空输入框">重新输入</span>' +
        '   <span class="search-btn-style search-btn-span click" title="点击搜索">🔍·搜索</span>' +
        '   <div class="clear"></div>' +
        '</div>' +
        //
        '<div class="search-btn-center do-btn-center must-btn select-none timeout-hide hide">' +
        '   <span class="search-btn-style copy-btn-span click" onclick="clear_copy(this)" data-clipboard-text="+86110">🎲·随机数</span>' +
        '   <span class="search-btn-style color-btn-span click">🌓·<span id="change-color-span"></span></span>' +

        '   <div class="clear"></div>' +
        '</div>' +
        //
        '<div class="swiper-container select-none">' +
        '    <div class="swiper-wrapper">' +

        //--
        '       <div class="swiper-slide more-btn">' +

        '           <div class="div-time select-none"></div>' +
        '           <div class="div-qr hide">' +
        '               <div class="div-qr-box" id="img-show_qr"></div>' +
        '           </div>' +

        '       </div>' +

        //--
        '       <div class="swiper-slide more-btn">' +

        '           <div class="search-btn-center quick-btn-center">' +
        '              <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="https://cn.investing.com/rates-bonds/china-10-year-bond-yield">CN十年国债</span>' +
        '              <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="https://cn.investing.com/rates-bonds/u.s.-10-year-bond-yield">US十年国债</span>' +
        '              <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="https://cn.investing.com/indices/volatility-s-p-500">VIX恐慌</span>' +

        '              <div class="clear"></div>' +
        '           </div>' +

        '           <div class="search-btn-center quick-btn-center">' +
        '              <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="https://cn.investing.com/commodities/crude-oil">WTI原油</span>' +
        '              <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="https://cn.investing.com/currencies/us-dollar-index">美元指数</span>' +
        '               <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="https://danjuanapp.com/valuation-table/jiucai">股债利差</span>' +

        '              <div class="clear"></div>' +
        '           </div>' +

        '           <div class="search-btn-center quick-btn-center ">' +
        '              <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="http://www.ywindex.com/">义乌指数</span>' +
        '              <span class="search-btn-style href-btn-span click" onclick="href_ext(this) " data-href="https://m.ximalaya.com/waiyu/44962493/">越语入门</span>' +
        '              <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="https://active.clewm.net/FwC95g?qrurl=http://qr35.cn/FwC95g&gtype=1&key=2fba516c7fe81e21f08824a04e2676d2f8043b2870">越语基础</span>' +

        '              <div class="clear"></div>' +
        '           </div>' +

        '           <div class="search-btn-center quick-btn-center ">' +
        '              <span class="search-btn-style href-btn-span click" onclick="href_ext(this) " data-href="https://learnku.com/docs/the-way-to-go/187-file/3736">Go入门指南L</span>' +
        '              <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="https://www.kancloud.cn/kancloud/the-way-to-go/81390">Go入门指南K</span>' +
        '              <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="https://learnku.com/docs/effective-go/2020/introduction/6236">Go高效编程</span>' +

        '              <div class="clear"></div>' +
        '           </div>' +

        '           <div class="search-btn-center quick-btn-center ">' +
        '              <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="http://c.biancheng.net/view/124.html">Go教程C</span>' +
        '              <span class="search-btn-style href-btn-span click" onclick="href_ext(this) " data-href="https://learnku.com/docs/gorm/v2/index/9728">GORM文档</span>' +
        '              <span class="search-btn-style href-btn-span click" onclick="href_ext(this) " data-href="https://learnku.com/docs/gorm/v2/index/9728">Gin文档</span>' +

        '              <div class="clear"></div>' +
        '           </div>' +

        '       </div>' +

        //--
        '       <div class="swiper-slide more-btn">' +

        '           <div class="search-btn-center quick-btn-center">' +
        '              <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="https://m.ithome.com">IT之家</span>' +
        '              <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="https://weibo.com/">微博</span>' +
        '               <span class="search-btn-style href-btn-span click" onclick="href_ext(this) " data-href="//t.bilibili.com">哔哩哔哩</span>' +
        '              <div class="clear"></div>' +
        '           </div>' +

        '           <div class="search-btn-center quick-btn-center ">' +
        '              <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="https://sspai.com/">少数派</span>' +
        '              <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="http://www.ruanyifeng.com/blog/">阮一峰</span>' +
        '              <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="https://www.v2ex.com/">V2EX</span>' +
        '              <div class="clear"></div>' +
        '           </div>' +

        '           <div class="search-btn-center quick-btn-center ">' +
        '              <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="https://www.huxiu.com/moment/">虎嗅7x24</span>' +
        '              <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="http://finance.sina.com.cn/7x24/">新浪7x24</span>' +
        '              <span class="search-btn-style href-btn-span click"  onclick="href_ext(this) " data-href="http://finance.sina.com.cn/stock/sl/#industry_1">新浪行业股</span>' +
        '              <div class="clear"></div>' +
        '           </div>' +

        '           <div class="search-btn-center quick-btn-center">' +


        '              <div class="clear"></div>' +
        '           </div>' +

        '       </div>' +

        //--
        '       <div class="swiper-slide more-btn hide">' +

        '           <div class="search-btn-center quick-btn-center">' +



        '              <div class="clear"></div>' +
        '           </div>' +
        '       </div>' +

        //--
        '    </div>' +

        '    <div class="swiper-pagination"></div>' +
        '    <div class="swiper-button-next hide"></div>' +
        '    <div class="swiper-button-prev hide"></div>' +

        '</div>' +
        //
        '<div class="clear"></div>';

    // 各种按钮事件绑定
    document.getElementsByClassName("input")[0].addEventListener("mouseover", function(e) {
        console_log("鼠标over了输入框，输入框自动聚焦");
        let that = this;
        that.focus();
    });
    document.getElementById("select").onchange = function(e) { // 设置历史和当前选中的搜索引擎
        console_log("选择搜素引擎");
        let that = this;
        set_search(that.value);
    };
    document.getElementById("input").onfocus = function(e) {
        console_log("监听输入框状态-onfocus");
        document.getElementsByClassName("select")[0].classList.add("liner-color");
        document.getElementById("input-div").classList.remove("input-div-blur");
        document.getElementById("input-div").classList.add("input-div-focus");
        dead_input(focus_time);
    };
    document.getElementById("input").onblur = function(e) {
        console_log("监听输入框状态-onblur");
        document.getElementsByClassName("select")[0].classList.remove("liner-color");
        document.getElementById("input-div").classList.remove("input-div-focus");
        document.getElementById("input-div").classList.add("input-div-blur");
        dead_input(blur_time);
    };

    //
    let db_click_time = (new Date()).getTime();
    document.onkeyup = function(event) { // Enter
        console_log("双击Enter进行搜素");
        let now_click_time = (new Date()).getTime();

        let _key = event.key;
        if (_key === "Enter") {
            let click_time = now_click_time * 1 - db_click_time * 1
            if (click_time <= 500) {
                db_click_time = 0;
                run_search();
            } else {
                db_click_time = (new Date()).getTime();
            }
        }
    };
    document.getElementsByClassName("search-btn-span")[0].addEventListener("click", function() {
        run_search();
    });
    document.getElementById("content-bg").addEventListener("click", function() {
        many_click(4, change_bg_color());
    });
    document.getElementsByClassName("color-btn-span")[0].addEventListener("click", function() {
        console_log("color-btn-span");
        change_bg_color();
    });
    document.getElementsByClassName("history-btn-span")[0].addEventListener("click", function() {
        clear_history();
        document.getElementById("input").value = "";
        setTimeout(function (){
            let now_url = window.location.href;
            window.location.replace(now_url);
        }, 100);
    });
    document.getElementsByClassName("refresh-btn-span")[0].addEventListener("click", function() {
        document.getElementById("input").value = "";
        // view.notice_txt("已清空输入框", 600);
        change_focus();
        // setTimeout(function (){
        //     let now_url = window.location.href;
        //     window.location.replace(now_url);
        // }, 100);
    });

    // 确定适应屏
    window.onresize = function() {
        let _width = window.innerWidth;
        let _screen_width = screen.width;
        resize_width(_width, _screen_width);
    };
    let width = window.innerWidth;
    let screen_width = screen.width;
    resize_width(width, screen_width);

    setTimeout(function() {
        try {
            document.getElementsByClassName("href-btn-center")[0].classList.add("hide");
        } catch (e) {
            console_log("跳过，-2");
        }
    }, 4000);

    setTimeout(function (){
        make_swiper();
    }, 100);

}

/*
*  处理历史记录
* */
function click_history(that){
    let data = that.getAttribute("data-history");
    document.getElementById("input").value = data;
    // run_search();
}
function show_history(){
    let data_key = "input_history";
    let array_key = "@=history=@";

    $("#input-history").html("");

    let data_string = view.get_data(data_key)
    // 限制历史记录长度
    let len = 10;
    let array_history = data_string.split(array_key)
    for (let i=0; i<len; i++){
       let the_history = array_history[i];
       if (the_history){
           let span = '<div class="history-span click select-none blue" data-history="'+the_history+'" onclick="click_history(this)" title="'+the_history+'" data-title="'+the_history+'">#'+the_history+'</div>'
           $("#input-history").append(span);
       }
    }
}
function update_history(){
    // let input_history = document.getElementById("input-history");
    let input = document.getElementById("input");

    let data_key = "input_history";
    let array_key = "@=history=@";
    let input_value = input.value;

    if (input_value){
        let data_string = view.get_data(data_key)
        // 去重历史记录
        if (view.string_include_string(data_string, input_value+"@=") !== -1){
            view.log("已存在历史记录：" + input_value);
        }else {
            // 限制历史记录长度
            let len = 30;
            let array_history = data_string.split(array_key)
            let new_data_string = "";
            for (let i=0; i<array_history.length; i++){
                let the_history = array_history[i];
                if (i<len){
                    new_data_string = the_history + array_key;
                    let new_data = input_value + array_key + data_string;
                    view.set_data(data_key, new_data)
                }
            }
        }

    }else {
        view.log("input_value不能为空")
    }

}
function clear_history(){
    $("#input-history").html("");
    let data_key = "input_history";
    return view.del_data(data_key);
}

// $(document).on("mouseenter", ".history-span", function (){
//     let that = $(this);
//     let title = that.data("title");
//     $("#input").val(title);
// });
//
// $(document).on("mouseleave", ".history-span", function (){
//     let that = $(this);
//     let title = that.data("title");
//     $("#input").val("");
// });

/*
 *  个性化颜色
 * */
const bg_cookie = search_cookie_pre + "bg_color";

function init_color() {
    let bg_color = getCookie(bg_cookie);
    if (bg_color === null || bg_color === ""){ // 默认颜色
        bg_color = 1;
    }else {
        bg_color = bg_color * 1;
    }

    let change_color_span = document.getElementById("change-color-span");
    let body = document.getElementsByClassName("body")[0];
    let select = document.getElementsByTagName("select")[0];
    let input = document.getElementsByTagName("input")[0];
    console_log("设置色：" + bg_color);

    if (bg_color === 0) { // 亮
        change_color_span.innerHTML = "雪白";

        body.classList.add("bg-light");
        body.classList.remove("bg-black");
        body.classList.remove("bg-yellow");
        body.classList.remove("bg-ivory");

        select.classList.add("select-color-light");
        input.classList.add("input-color-light");
        select.classList.remove("select-color-black");
        input.classList.remove("input-color-black");
        select.classList.remove("select-color-yellow");
        input.classList.remove("input-color-yellow");
        select.classList.remove("select-color-ivory");
        input.classList.remove("input-color-ivory");
    } else if (bg_color === 1) { // 暗
        change_color_span.innerHTML = "昏黑";

        body.classList.remove("bg-light");
        body.classList.add("bg-black");
        body.classList.remove("bg-yellow");
        body.classList.remove("bg-ivory");

        select.classList.remove("select-color-light");
        input.classList.remove("input-color-light");
        select.classList.add("select-color-black");
        input.classList.add("input-color-black");
        select.classList.remove("select-color-yellow");
        input.classList.remove("input-color-yellow");
        select.classList.remove("select-color-ivory");
        input.classList.remove("input-color-ivory");
    } else if (bg_color === 2) { // 黄
        change_color_span.innerHTML = "夕黄";

        body.classList.remove("bg-light");
        body.classList.remove("bg-black");
        body.classList.add("bg-yellow");
        body.classList.remove("bg-ivory");

        select.classList.remove("select-color-light");
        input.classList.remove("input-color-light");
        select.classList.remove("select-color-black");
        input.classList.remove("input-color-black");
        select.classList.add("select-color-yellow");
        input.classList.add("input-color-yellow");
        select.classList.remove("select-color-ivory");
        input.classList.remove("input-color-ivory");
    } else if (bg_color === 3) { // 象牙
        change_color_span.innerHTML = "牙白";

        body.classList.remove("bg-light");
        body.classList.remove("bg-black");
        body.classList.remove("bg-yellow");
        body.classList.add("bg-ivory");

        select.classList.remove("select-color-light");
        input.classList.remove("input-color-light");
        select.classList.remove("select-color-black");
        input.classList.remove("input-color-black");
        select.classList.remove("select-color-yellow");
        input.classList.remove("input-color-yellow");
        select.classList.add("select-color-ivory");
        input.classList.add("input-color-ivory");
    } else { // 默认
        view.alert_txt("无默认颜色，无法渲染")
    }
}

function change_bg_color() {
    let bg_color = getCookie(bg_cookie);
    if (bg_color === null || bg_color === ""){ // 默认颜色
        bg_color = 1;
    }else {
        bg_color = bg_color * 1;
    }

    // 0=bg-light；1=bg-black；2=bg-yellow；
    if (bg_color === 0) { // 切换到第二个
        setCookie(bg_cookie, (bg_color + 1));
    } else if (bg_color === 1) { // 切换到第三个
        setCookie(bg_cookie, (bg_color + 1));
    } else if (bg_color === 2) { // 切换到第四个
        setCookie(bg_cookie, (bg_color + 1));
    } else if (bg_color === (bg_color + 1)) { //  // 切换到第一个
        setCookie(bg_cookie, 0);
    } else { // 默认为0
        setCookie(bg_cookie, 0);
    }

    init_color();
}

function change_focus() {
    document.getElementById("input").focus();
}

function change_blur() {
    document.getElementById("input").blur();
}

// 调整屏幕宽度变化时的页面展示适应性
function resize_width(width, screen_width) {
    console_log(width);

    // if (width < 400){
    //     document.getElementsByClassName("body")[0].classList.add("zoom-80");
    // }else {
    //     document.getElementsByClassName("body")[0].classList.remove("zoom-80");
    // }

    try {
        if (width < 1200) {
            document.getElementsByClassName("href-btn-center")[0].classList.add("hide");
        } else {
            document.getElementsByClassName("href-btn-center")[0].classList.remove("hide");
            setTimeout(function() {
                document.getElementsByClassName("href-btn-center")[0].classList.add("hide");
            }, 4000);
        }
    } catch (e) {
        console_log("跳过，-1");
    }


    if (screen_width < 600) {
        let img = "";
        try {
            img = document.getElementsByClassName("qr_img")[0].getAttribute("src");
        } catch (e) {
            img = "";
        }
        if (img) {
            document.getElementsByClassName("div-qr")[0].classList.remove("hide");
        } else {
            console_log("img为空则跳过");
        }
    } else {
        document.getElementsByClassName("div-qr")[0].classList.add("hide");
    }

}

/*
 * 提醒
 * */
function write_tips_text(text) {
    document.getElementsByClassName("res-div")[0].innerHTML = '<div class="flex-center tips-div select-none hide" ><div class="res-txt">' + text + '</div></div>';
}

function show_loading() {
    console_log("展示遮蔽层");
    document.getElementById("loading-div").classList.remove("hide");
}

function delete_loading() {
    console_log("删除遮蔽层");
    document.getElementById("loading-div").classList.add("hide");
}

function href_ext(that) {
    let el_href = that.getAttribute("data-href");
    console_log(el_href);

    if (el_href) {
        window.open(el_href, "_blank");
    } else {
        view.alert_txt("参数不能为空", 2000);
        console_log("参数不能为空");
    }
}

// 时间
function timer() {
    return;
    try {
        // 报时
        if (view.get_date()[11] == "00" && view.get_date()[12] == "01"){
            let txt = view.get_date()[10]+"点整";
            view.notice_txt(txt, 5000);
            view.voice(txt, 0.7, false)
        }
        // view.get_date()[11] == "30" &&
        if (view.get_date()[11] == "30" && view.get_date()[12] == "01"){
            let txt = view.get_date()[10]+"点"+view.get_date()[11]+"分";
            view.notice_txt(txt, 5000);
            view.voice(txt, 0.5, false)
        }

        // dom
        document.getElementsByClassName('div-time')[0].innerHTML =
            '<div class="div-time-hour">' +
            "  " + view.get_date()[7] + // 时刻
            '</div>' +
            '<div class="div-time-date">' +
            '   <span class="div-time-span">' + view.get_date()[2] + '/' + view.get_date()[5] + '</span>' +
            // '   <span class="div-time-span">' + view.get_date()[9] + '</span>' +
            '   <span class="div-time-span">' + view.get_date()[9] + '</span>' +
            '</div>' +
            "";
    } catch (e) {
        view.log("跳过");
    }
}

//
function make_swiper(){
    let swiper = new Swiper('.swiper-container', {
        autoHeight: true, //enable auto height
        spaceBetween: 200,
        //loop: true,
        pagination: {
            el: '.swiper-pagination',
            clickable: true,
            renderBullet: function (index, className) { // 底部数字
                return '<span class="' + className + '">' + (index + 1) + '</span>';
            },
        },
        keyboard: { // 键盘方向键控制
            enabled: true,
        },
        mousewheel: true, // 鼠标滚轮控制
        navigation: {
            nextEl: '.swiper-button-next',
            prevEl: '.swiper-button-prev',
        },
    });

    setTimeout(function (){
        timer();
        setInterval(function() {
            timer();
        }, 1000);
    }, 100);

    setTimeout(function (){
        $(".swiper-pagination").addClass("hide");
        $(".timeout-hide").css("opacity", "0");
    }, 10);

}

//
function clear_copy(that){
    // view.notice_txt(view.read_clipboard());

    let clipboard = new Clipboard(".copy-btn-span");
    clipboard.on('success', function(e) {
        // console.info('Action:', e.action);
        // console.info('Text:', e.text);
        // console.info('Trigger:', e.trigger);
        view.alert_txt("已生成随机数到粘贴板", 1200);
        e.clearSelection();
    });
    clipboard.on('error', function(e) {
        // console.error('Action:', e.action);
        // console.error('Trigger:', e.trigger);
        view.alert_txt("粘贴板操作失败", 800);
    });

}

function start_this_page(info) {
    view.log(info);
    // view.log("主框架解析完成，开始渲染模块页面 > >");

    init_dom();
    init_color();
    delete_loading();

    show_history();

}