
// 推出所有设备
function all_layout(){
    view.alert_txt("正在退出全部设备..", "long", "clear");

    /*开始-请求数据*/
    $.ajax({
        url: api_url + "admin/admin_all_layout",
        type: "POST",
        dataType: "json",
        async: true,
        data: { // 字典数据
            app_class: app_class,
            login_name: login_name,
            login_token: login_token,
            login_id: login_id,
            login_level: login_level,
        },
        success: function(back, status){
            // 数据转换为json
            let data = "";
            let text = "";
            if(typeof back === "string"){
                data = JSON.parse(back);
                text = back;
            } else {
                data = back;
                text = JSON.stringify(back);
            }
            // 校验数据规格
            if (typeof data.state !== "undefined" && typeof data.msg !== "undefined"){
                // 解析json
                if (data.state === 0){ // 无数据或参数不全
                    view.alert_txt(data.msg, 2000, "clear");
                }else if (data.state === 1){ // 接口数据成功
                    view.alert_txt(data.msg, 500, "clear");
                    must_login();
                }else if (data.state === 302){ // 需要重新登录或初始化数据
                    must_login(data.msg);
                }else if (data.state === 403){ // 拒绝访问
                    must_login(data.msg);
                }else {
                    let txt = data.msg+"("+ data.state +")";
                    console.error(txt);
                    must_login(data.msg);
                }
            }else {
                let info = "类型：" + typeof back + "\n数据：" + text +"\n状态：" + status + "。";
                view.error("=接口返回未知规格的参数=\n" + info);
            }
        },
        error: function (xhr) {
            console.error(xhr);
        }
    });
    /*结束-请求数据*/

}

// 退出本设备
function that_layout(){
    view.alert_txt("正在退出本设备..", "long", "clear");
    must_login();
}

// 跳转到登录页
function must_login(msg){
    if (!msg){msg="请先登录..";}
    view.alert_txt(msg, "long");
    view.del_cookie("login_token");
    view.del_cookie("login_id");
    view.del_cookie("login_level");

    setTimeout(function (){
        window.location.replace("./?route=login&back_url=" + encodeURIComponent(window.location.href));
    }, 500);
}

//
function level_dom(level){
    // level-dom hide level-2 level-3
    let level_hide_txt = "";
    if (level === 1){
        $(".level-dom").removeClass("hide");
        $(".level-"+level+"-hide").html(level_hide_txt).addClass("hide");
    }else if (level === 2){
        $(".level-dom").removeClass("hide");
        $(".level-"+level+"-hide").html(level_hide_txt).addClass("hide");
    }else if (level === 3){
        $(".level-dom").removeClass("hide");
        $(".level-"+level+"-hide").html(level_hide_txt).addClass("hide");
    } else if (level === 4){
        $(".level-dom").removeClass("hide");
        $(".level-"+level+"-hide").html(level_hide_txt).addClass("hide");
    }else {
        $(".level-dom").html("<h3>null-level</h3>");
    }

}

// 开发状态时dom展示，生产环境时dom隐藏。class=show_state_dev
function show_state(level){
    // 判断是否是生产状态
    let host = window.location.host;
    for (let i=0; i<app_domain.length; i++){
        let the_domain = app_domain[i];
        //view.log([the_domain, host]);
        if (the_domain){
            if (view.string_include_string(host, the_domain) !== -1){
                //
                $(".show-state-dev").addClass("hide");
                //
                break;
            }
        }
    }
}

//
// 检查admin_token
function check_admin_token(){

    /*开始-请求数据*/
    $.ajax({
        url: api_url + "admin/check_admin_login",
        type: "POST",
        dataType: "json",
        async: true,
        data: { // 字典数据
            app_class: app_class,
            url: window.location.href,
            login_name: login_name,
            login_token: login_token,
            login_id: login_id,
            login_level: login_level,
        },
        success: function(back, status){
            // 数据转换为json
            let data = "";
            let text = "";
            if(typeof back === "string"){
                data = JSON.parse(back);
                text = back;
            } else {
                data = back;
                text = JSON.stringify(back);
            }
            // 校验数据规格
            if (typeof data.state !== "undefined" && typeof data.msg !== "undefined"){
                // 解析json
                if (data.state === 0){ // 无数据或参数不全
                    view.alert_txt(data.msg, 2000, "clear");
                }else if (data.state === 1){ // 接口数据成功
                    let the_level = data.content.level;
                    let the_level_name = data.content.level_name;
                    view.log([the_level, the_level_name]);

                    login_level_name = the_level_name;

                    // 1超级管理员，2组管理员，3普通管理员，4仅查看数据
                    level_dom(the_level);
                    show_state(the_level);
                    if (the_level === 1){
                        start_this_page(the_level);
                        try {start_foot();}catch (e){view.error(e);}
                    }
                    else if (the_level === 2){
                        start_this_page(the_level);
                        try {start_foot();}catch (e){view.error(e);}
                    }
                    else if (the_level === 3){
                        start_this_page(the_level);
                        try {start_foot();}catch (e){view.error(e);}
                    }
                    else if (the_level === 4){
                        start_this_page(the_level);
                        try {start_foot();}catch (e){view.error(e);}
                    }
                    else {
                        view.alert_txt(data.msg, "long", "clear");
                    }

                    // 其他
                    $(".mine-login_name").html(login_level_name + "（"+login_nickname+"）");

                }else if (data.state === 302){ // 需要重新登录或初始化数据
                    must_login(data.msg);
                }else if (data.state === 403){ // 拒绝访问
                    must_login(data.msg);
                }else {
                    let txt = data.msg+"("+ data.state +")";
                    console.error(txt);
                    must_login(data.msg);
                }
            }else {
                let info = "类型：" + typeof back + "\n数据：" + text +"\n状态：" + status + "。";
                view.error("=接口返回未知规格的参数=\n" + info);
            }
        },
        error: function (xhr) {
            console.error(xhr);
            view.alert_txt("接口请求错误或访问被限制", 2500, "clear");
        }
    });
    /*结束-请求数据*/

}



// 页面必要安全参数校验
function must_safe_check(e){
    if(!navigator.webdriver){

        if (debug === false && view.is_user_screen() === false){ // 生产状态
            $(".depend").addClass("hide");
            view.alert_txt("不符合生产状态下的设备要求，页面停止解析。", "long", "clear");
            return;
        }

        let route = view.get_url_param("", "route");
        if (route === "login"){
            start_this_page(e);
        } else { // 不是login的话就直接检查是否已经登录
            start_this_page(e);
        }
    }else {
        view.alert_txt("请不要使用模拟浏览器操作页面信息！", "long", "clear");
    }

}
