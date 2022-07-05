

(function () { // 推荐组件里面的函数写在闭包里面

    let part = {
        "part_init": function (e) {
            view.log(e);

            // 渲染选中
            let active_div = view.get_url_param("", "route");
            $(".active-" + active_div).addClass("active-div");

        },

    };

    // 写入组件
    view.write_htm(page_url + "parts/nav/nav.html", "depend", function () {
        part.part_init("nav");
    });
    view.write_css([page_url + "parts/nav/nav.css"], function () {
        view.log("css");
    });

})();

