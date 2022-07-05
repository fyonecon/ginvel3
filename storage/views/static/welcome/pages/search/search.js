

function jump_location(engine, word, url) {
    let page_time = view.get_date()[0] + "_" + view.md5(view.get_date()[1]);
    let _word = '';

    if (word === "%s" || word === ""){
        view.alert_txt("word参数出问题或搜索内容不能为空", 2000);
        url = "https://www.bing.com/?ensearch=1&q=如何使用搜索";
    }
    else if (word === "首页" || word === "home" || word === "ggvs"){
        url = "./";
    }
    else {
        //let del_fake_news = " -aliyun.com -huaweicloud.com"
        let del_fake_news = " "
        del_fake_news = decodeURI(del_fake_news)

        try {
            _word = decodeURI(word);
        }catch (e) {
            _word = word;
        }

        if (engine === "baidu"){
            url = "https://www.baidu.com/s?ie=utf-8";
            url = url + "&wd=" + _word + del_fake_news + "&page_time=" + page_time;
        }else if (engine === "bing"){
            url = "https://www.bing.com/?ensearch=1";
            url = url + "&q=" + _word + "&page_time=" + page_time;
        }
        else if (engine === "google"){
            url = "https://www.google.com/search?q=";
            url = url + _word + "&page_time=" + page_time;
        }
        else if (engine === "m-toutiao"){
            url = "https://m.toutiao.com/search/?keyword=";
            url = url + _word + "&page_time=" + page_time;
        }
        else if (engine === "toutiao"){
            url = "https://www.toutiao.com/search/?keyword=";
            url = url + _word + "&page_time=" + page_time;
        }
        else if (engine === "music"){
            url = "https://www.hifini.com/search-";
            url = url + _word + "-1-1-1.htm?page_time=" + page_time;
        }
        else if (engine === "video"){
            url = "https://www.bing.com/search?ensearch=1&q=tokyvideo+";
            url = url + _word + "&page_time=" + page_time;
        }
        else {
            view.alert_txt("engine参数为空，不能选择跳转的目标地址");
            view.log("/?route=search&engine=&word=");
            return;
        }

    }

    setTimeout(function () {
        window.location.replace(url);
    }, 500);
}


function jump_to_search_engine() {
    let engine = ""; // 哪个搜索引擎
    let word = ""; // 搜索的关键词
    let url = "";

    engine = view.get_url_param("", "engine");
    try {
        word = view.get_url_param("", "word");
    }catch (e) {
        view.error(["可忽略的错误", e]);
        word = "";
    }

    if (!engine){
        let search_eq = view.get_cookie("search__eq");
        search_eq = 1*search_eq;

        if (search_eq === 0){
            engine = "bing";
        }else if (search_eq === 1){
            engine = "google";
        }else if (search_eq === 2){
            engine = "baidu";
        }else if (search_eq === 3){
            engine = "toutiao";
        }else {
            engine = "bing";
        }

        view.log([word, engine, search_eq]);
    }else {
        view.log([engine]);
    }

    jump_location(engine, word, url);
}


function start_this_page(e) {
    view.log(e);
    jump_to_search_engine();
}
