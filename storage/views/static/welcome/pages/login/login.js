
function clear() {
	$(".user-login-name").val("");
	$(".user-login-pwd").val("");
	$(".user-input-check").val("");
}

function admin_login() {
	let back_url = view.get_url_param("", "back_url");

	let name = $(".user-login-name").val().trim();
	let pwd = $(".user-login-pwd").val().trim();
	let num = $(".user-input-check").val().trim()*1;

	if (!name){
		view.notice_txt("登录名未填写", 2000);
		return;
	}
	if (!pwd){
		view.notice_txt("登录密码未填写", 2000);
		return;
	}else{
		pwd = view.md5(pwd);
	}

	if (!num || num_res !== num){
		view.notice_txt("验证码计算结果不正确", 2000);
		make_num();
		return;
	}

    var inner_w = window.innerWidth; if (!inner_w){inner_w=0;}
    var inner_h = window.innerHeight; if (!inner_h){inner_h=0;}
    var screen_w = window.screen.width; if (!screen_w){screen_w=0;}
    var screen_h = window.screen.height; if (!screen_h){screen_h=0;}
    var screen_h_foot = screen_h - inner_h;
    let _screen = screen_w + "x" + screen_h + "_" + inner_w + "x" + inner_h + "=" + screen_h_foot;
    var _ua = window.navigator.userAgent.toLowerCase();

	view.alert_txt("正在登录..", "long");

	/*请求数据*/
	$.ajax({
		url: api_url+"admin/admin_login",
		type: "POST",
		dataType: "json",
		async: true,
		data: { // 字典数据
            app_token: app_token,
            app_uid: app_uid,
			login_name: name,
			login_pwd: pwd,
            screen: _screen,
            ua: _ua,
		},
		success: function(back, status){

			// 数据转换为json
			let res;
			let data = "";
			if(typeof back === "string"){
				res = JSON.parse(back);
				data = back;
			} else {
				res= back;
				data = JSON.stringify(back)
			}
			view.log("状态：" + status +"\n类型：" + typeof back + "。\n数据：" + data );

			make_num();
			clear();

			// 解析json
			if (res.state === 0){
				view.alert_txt(res.msg, 2000, "clear");

			}else if (res.state === 1) {
				view.log(res.msg);
				view.alert_txt(res.msg, 3000);

				let login_token = res.content.login_token;
				let login_id = res.content.login_id;
				let login_level = res.content.login_level;
				let login_name = res.content.login_name;
				let login_nickname = res.content.login_nickname;

				let href = res.content.jump_url;

				let time = 30*24*60*60*1000; // 30天过期

				view.set_cookie("login_name", login_name, time);
				view.set_cookie("login_token", login_token, time);
				view.set_cookie("login_id", login_id, time);
				view.set_cookie("login_level", login_level, time);
				view.set_cookie("login_nickname", login_nickname, time);

				setTimeout(function () {
					if (back_url){
						back_url = decodeURIComponent(back_url);
						href = back_url;
					}else {
						if (!href){
							href = "./";
						}
					}
					view.log(href);
					window.location.replace(href);
				}, 500);

			}else if (res.state === 2){
				view.alert_txt(res.msg, 2000, "clear");
			}else {
				view.alert_txt("超范围的state", 3000, "clear");
			}
		},
		error: function (xhr) {
			console.log(xhr);
			view.alert_txt("接口请求错误或者网络不通", 2500, "clear");
		}
	});

}

function app_action(){ // 每次重新生成即可，不限次数，但限制频率
    let app_uid_info = view.make_app_uid(app_class);

    let screen = window.screen.width + "x" + window.screen.height + "_" + window.innerWidth + "x" + window.innerHeight;
    let refer = document.referrer;
    let href = window.location.href;
    let _app_uid = app_uid_info[0];
    let _app_date = app_uid_info[1];
    let ua = window.navigator.userAgent.toLowerCase();

    view.log([app_class, screen, refer, href, _app_uid, _app_date, ua]);

    /*请求数据*/
    $.ajax({
        url: api_url+"app/app_action",
        type: "POST",
        dataType: "json",
        async: true,
        data: { // 字典数据
            app_class: app_class,
            app_name: app_name,
            screen: screen,
            refer: refer,
            href: href,
            app_uid: _app_uid,
            app_date: _app_date,
            ua: ua,
        },
        success: function(back, status){
            // 数据转换为json
            let res;
            let data = "";
            if(typeof back === "string"){
                res = JSON.parse(back);
                data = back;
            } else {
                res= back;
                data = JSON.stringify(back)
            }
            // 解析json
            if (res.state === 0){
                view.alert_txt(res.msg, 3000, "clear");
            }else if (res.state === 1) {
                let info = res.content;

                let the_app_token = info.app_token;
                let the_app_uid = info.app_uid;

                let time = 2*60*60*1000; // 2天过期

                view.set_cookie("app_token", the_app_token, time);
                view.set_cookie("app_uid", the_app_uid, time);

                app_token = the_app_token;
                app_uid = the_app_uid;

                // 其他
                $(".user-title-span").html(app_name);
                make_num();

            }else if (res.state === 2){
                view.alert_txt(res.msg, 3000, "clear");
            }else if (res.state === 301){
                let href_301 = res.content.href_301;
                window.location.replace(href_301);
            }else {
                view.alert_txt("超范围的state", 3000, "clear");
            }
        },
        error: function (xhr) {
            console.log(xhr);
            view.alert_txt("接口请求错误或被限制刷新频率", 3000, "clear");
        }
    });

}

function make_num() {

	let num1= view.js_rand(50, 99);
	let num2= view.js_rand(1, 49);

	let num0 = view.js_rand(1, 3); // 加减乘除
	let mark = "";

	switch (num0) {
		case 1:
			num_res = num1-num2;
			mark = "减";
			break;
		case 2:
			num_res = num1+num2;
			mark = "加";
			break;
		case 3:
			num_res = num1*num2;
			mark = "乘";
			break;
		default:
			num_res = num1*num2;
			mark = "乘";
			break;
	}

	let div = num1+"&nbsp;"+mark+"&nbsp;"+num2+"&nbsp;"+"&nbsp;"+"等于";
	$(".show-check").html(div);
}

(function () {

	$(document).on("click", ".user-login-btn", function () {
        if (login_state === 1){
            admin_login();
        }else {
            view.notice_txt("初始化失败", 2000);
        }
	});

	document.onkeydown = function(e){
		let ev = document.all ? window.event : e;
		if(ev.keyCode==13) {
            if (login_state === 1){
                admin_login();
            }else {
                view.notice_txt("初始化失败", 2000);
            }
		}
	};

	$(document).on("click", ".show-check", function () {
        if (login_state === 1){
            make_num();
        }else {
            view.notice_txt("初始化失败", 2000);
        }
	});

})();


// 页面数据入口
let num_res = 0;
let login_state = 0;
function start_this_page() {
    $(".user-login-name").val("576");
    $(".user-login-pwd").val("789");
    login_state = 1;
    app_action();
}
