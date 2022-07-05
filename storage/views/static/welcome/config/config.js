/*自定义配置页面的一些全局参数*/

/*
* 1）页面生命周期（index.html--config.js--框架解析index.js--公共all.js/css文件--执行wifi广告劫持严重和清除--pages.htm--pages.js--page_loaded.js）。
* 2）不依赖于node但需依赖服务端环境，或者依赖CDN环境。
* */

const debug = true;                     // 调试模式，统一打印日志，true & false

// 框架渲染的必要参数
const cookie_prefix   = "view_ggvs_";   // cookie前缀
const route_404       = "?route=404";   // 404
const route_default   = "?route=home";  // 页面进入默认页

const file_url        = index_file_url?index_file_url:"./";             // 资源文件CDN主域名（js、css、img、font等资源文件）//cdnaliyun.oss-cn-hangzhou.aliyuncs.com/view-ggvs/
const page_url        = index_file_url?index_file_url:"./";             // htm文件的服务器地址，因为使用了ajax请求，不能直接请求本地文件，可以全部放在CDN里面

const cache_time      = 1000000;            // 缓存时间：ms
const page_time       =  "" + Math.floor((new Date()).getTime()/cache_time)*cache_time;

const api_url         = "https://xxx.com/api/";  // api主地址
const page_title      = " - ggvs";

// 白名单refer域名
let refer = [
    {
        'check_refer': false, // 是否开启白名单refer检测
        'jump_site': '', // 遇到黑名单refer的落地地址
        'white_refer': [

        ],
    }
];

// 自定义
// App验证参数
let app_class = "view_ggvs";
let app_version = "v3.2.0";
let app_name = "view单页web框架";
let app_platform = "WebApp";

// 登录用户使用的验证参数
let login_token = "";
let login_id = 0;
let login_name = "";
let login_level = 0;
let login_level_name = "（未知等级）";
let login_nickname = "（未登录）";



