/*参与所有组件中*/
/*组件加载完，和组件js权重差不多，比all.js载入时间迟*/
/*可以定义一些组件成功载入完成后的一些【个性化东西】*/

// 组件加载时间
view.log("框架解析用时=" + (time_loaded - time_start) + "ms");
view.log("框架报错时间=" + time_error);

// 获取客户端IP
// view.write_js(["http://pv.sohu.com/cityjson?ie=utf-8"], function () {
//     let cip = returnCitySN["cip"];
//     let cname = returnCitySN["cname"];
//     let cid = returnCitySN["cid"];
//
//     view.log([cip, cname, cid]);
//
// });

(function (){

    let the_route = view.get_url_param("", "route");

    window.onload = function (){
        //

    };

    window.onresize = function (){
        let width = window.innerWidth;
        //
    };

    let before_scroll = 0;
    window.onscroll = function() {
        let scroll = document.documentElement.scrollTop || document.body.scrollTop;
        let direct = scroll-before_scroll;

        //
        if (direct > 5 && direct < 80){ // 向下滚动
            //
        }else if (direct < -5 && direct > -80){ // 向上滚动
            //
        }else {
            //
        }
        before_scroll = scroll;

        // 隐藏顶部状态栏
        let width = window.innerWidth;
        if (width>=780){
            // if (scroll > 100){
            //     $(".header-box").fadeOut(500);
            // }else {
            //     $(".header-box").fadeIn(500);
            // }
        }else {
            view.log("跳过");
        }

        // 固定tab切换部分
        try {
            let tab_top = $(".tab-btn-box").offset().top;
            let scroll_top = $(window).scrollTop();
            //view.log(["tab_top=", tab_top, "scroll_top=", scroll_top, "tab_top-scroll_top=", tab_top-scroll_top]);
            if (tab_top-scroll_top < 60){
                $(".tab-btn-box-nav").addClass("tab-btn-box-nav-fixed");
            }else {
                $(".tab-btn-box-nav").removeClass("tab-btn-box-nav-fixed");
            }
        }catch (e) {}


    }


    $(document).on("click", ".header-menu", function (){
        let that = $(this);

        if (!that.hasClass("transform1-90")){ // 展示
            that.removeClass("transform1-0");
            that.addClass("transform1-90");

            $(".header-bg").removeClass("hide");
            setTimeout(function (){
                $(".header-menu-list").slideDown(200);
            }, 200);
        }else { // 隐藏
            that.removeClass("transform1-90");
            that.addClass("transform1-0");

            $(".header-menu-list").slideUp(200);
            setTimeout(function (){
                $(".header-bg").addClass("hide");
            }, 200);
        }

    });


    $(document).on("click", ".header-bg", function (){
        let that = $(this);
        $(".header-menu").click();
    });

    $(document).on("click", ".a-fresh", function (){
        let that = $(this);
        window.location.reload();
    });

    $(document).on("click", ".direct-home", function (){
        let that = $(this);
        window.location.href = "./?route=home";
    });


    $(".input-jump_url_certi").addClass("hover");
    $(document).on("dblclick", ".input-jump_url_certi", function (){
        let that = $(this);
        let href = that.val();
        if (href){
            window.open(href, "_blank");
        }else {
            view.notice_txt("链接不能为空", 2000);
        }
    });
    $(".input-jd_url").addClass("hover");
    $(document).on("dblclick", ".input-jd_url", function (){
        let that = $(this);
        let href = that.val();
        if (href){
            window.open(href, "_blank");
        }else {
            view.notice_txt("链接不能为空", 2000);
        }
    });

    // js 监听页面应用进入后台
    document.addEventListener("visibilitychange",function(){
        //
        if(document.visibilityState == "visible"){
            view.log("切换到前台："+view.get_date()[0]);
            try {show_this_page(["可选初始切换到前台函数"]);}catch (e){view.log("show_this_page()：可忽略的可选初始切换到前台函数");}
        }
        //
        if(document.visibilityState == "hidden"){
            view.log("切换到后台："+view.get_date()[0]);
            try {hide_this_page(["可选初始切换到后台函数"]);}catch (e){view.log("hide_this_page()：可忽略的可选初始切换到后台函数");}
        }
    });

    //tab
    $(document).on("click", ".tab-btn-item", function (){
        let that = $(this);
        that.parent(".tab-btn-box-nav").parent(".tab-btn-box").find(".tab-btn-item").removeClass("active-tab-btn");
        that.addClass("active-tab-btn");
        let item = that.data("item");
        let item_active = that.data("item-active");
        view.log([item, item_active]);
        $("."+item).addClass("hide");
        $("."+item_active).removeClass("hide");
    });

    //
    // if (window.screen.width > 280 && window.screen.width <= 640){
    //     $(".body").css({"zoom":"0.80",});
    // }
    // else if (window.screen.width > 640 && window.screen.width <= 1366){
    //     $(".body").css({"zoom":"0.85",});
    // }
    // else if (window.screen.width > 1366 && window.screen.width <= 2660){
    //     $(".body").css({"zoom":"0.90",});
    // }
    // else {
    //     $(".body").css({"zoom":"0.95",});
    // }

})();






