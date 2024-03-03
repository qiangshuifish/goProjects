package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestErr(t *testing.T) {
	html := "<!DOCTYPE html> <html> <head> <meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" /> <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title> </body> </html>"
	xpath := ""
	//config := "{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\",\"obj\":{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\"}}"
	srt := ""
	fmt.Println(Exe(srt, xpath, html, ""))
}

func TestXpath(t *testing.T) {
	html := "<!DOCTYPE html> <html> <head> <meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" /> <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title> </body> </html>"
	xpath := "//title"
	//config := "{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\",\"obj\":{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\"}}"
	srt := ""
	fmt.Println(Exe(srt, xpath, html, ""))
}

func TestXpathAndHandle(t *testing.T) {
	html := "<!DOCTYPE html> <html> <head> <meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" /> <title id=\"1\">asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title id=\"2\">asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title id=\"3\">asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title id=\"4\">asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title >asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title> </body> </html>"
	xpath := "//title"
	//config := "{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\",\"obj\":{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\"}}"
	srt := ""
	resultHandle := `{"value_content":"//title","id_attr":"//title"}`
	fmt.Println(Exe(srt, xpath, html, resultHandle))
}

func TestXpathAndHandle02(t *testing.T) {
	html := "<tr>\\n <td class=\\\"pls\\\" rowspan=\\\"2\\\">\\n<div id=\\\"favatar19133954\\\" class=\\\"pls cl favatar\\\">\\n<div class=\\\"pi\\\">\\n<div class=\\\"authi\\\"><a href=\\\"space-uid-446740.html\\\" target=\\\"_blank\\\" class=\\\"xw1\\\">blue2023</a>\\n</div>\\n</div>\\n<div class=\\\"p_pop blk bui card_gender_0\\\" id=\\\"userinfo19133954\\\" style=\\\"display: none; margin-top: -11px;\\\">\\n<div class=\\\"m z\\\">\\n<div id=\\\"userinfo19133954_ma\\\"></div>\\n</div>\\n<div class=\\\"i y\\\">\\n<div>\\n<strong><a href=\\\"space-uid-446740.html\\\" target=\\\"_blank\\\" class=\\\"xi2\\\">blue2023</a></strong>\\n<em>当前离线</em>\\n</div><dl class=\\\"cl\\\">\\n<dt>积分</dt><dd><a href=\\\"home.php?mod=space&amp;uid=446740&amp;do=profile\\\" target=\\\"_blank\\\" class=\\\"xi2\\\">259</a></dd>\\n</dl><div class=\\\"imicn\\\">\\n<a href=\\\"home.php?mod=space&amp;uid=446740&amp;do=profile\\\" target=\\\"_blank\\\" title=\\\"查看详细资料\\\"><img src=\\\"static/image/common/userinfo.gif\\\" alt=\\\"查看详细资料\\\"/></a>\\n</div>\\n<div id=\\\"avatarfeed\\\"><span id=\\\"threadsortswait\\\"></span></div>\\n</div>\\n</div>\\n<div>\\n<div class=\\\"avatar\\\" onmouseover=\\\"showauthor(this, &#39;userinfo19133954&#39;)\\\"><a href=\\\"space-uid-446740.html\\\" class=\\\"avtm\\\" target=\\\"_blank\\\"><img src=\\\"https://qz1gy.app/uc_server/data/avatar/000/44/67/40_avatar_middle.jpg\\\" onerror=\\\"this.onerror=null;this.src=&#39;https://qz1gy.app/uc_server/images/noavatar_middle.gif&#39;\\\"/></a></div>\\n</div>\\n<div class=\\\"tns xg2\\\"><table cellspacing=\\\"0\\\" cellpadding=\\\"0\\\"><tbody><tr><th><p><a href=\\\"home.php?mod=space&amp;uid=446740&amp;do=profile\\\" class=\\\"xi2\\\">724</a></p>金钱</th><th><p><a href=\\\"home.php?mod=space&amp;uid=446740&amp;do=profile\\\" class=\\\"xi2\\\">0</a></p>色币</th><td><p><a href=\\\"home.php?mod=space&amp;uid=446740&amp;do=profile\\\" class=\\\"xi2\\\">9</a></p>评分</td></tr></tbody></table></div>\\n\\n<p><em><a href=\\\"home.php?mod=spacecp&amp;ac=usergroup&amp;gid=15\\\" target=\\\"_blank\\\">Lv6 后起之秀</a></em></p>\\n\\n\\n<p><span class=\\\"pbg2\\\" id=\\\"upgradeprogress_19133954\\\" onmouseover=\\\"showMenu({&#39;ctrlid&#39;:this.id, &#39;pos&#39;:&#39;12!&#39;, &#39;menuid&#39;:&#39;g_up19133954_menu&#39;});\\\"><span class=\\\"pbr2\\\" style=\\\"width:19%;\\\"></span></span></p>\\n<div id=\\\"g_up19133954_menu\\\" class=\\\"tip tip_4\\\" style=\\\"display: none;\\\"><div class=\\\"tip_horn\\\"></div><div class=\\\"tip_c\\\">Lv6 后起之秀, 积分 259, 距离下一级还需 241 积分</div></div>\\n\\n<dl class=\\\"pil cl\\\">\\n  <dt>积分</dt><dd><a href=\\\"home.php?mod=space&amp;uid=446740&amp;do=profile\\\" target=\\\"_blank\\\" class=\\\"xi2\\\">259</a></dd>\\n</dl>\\n\\n\\n<ul class=\\\"xl xl2 o cl\\\">\\n<li class=\\\"pm2\\\"><a href=\\\"home.php?mod=spacecp&amp;ac=pm&amp;op=showmsg&amp;handlekey=showmsg_446740&amp;touid=446740&amp;pmid=0&amp;daterange=2&amp;pid=19133954&amp;tid=1901572\\\" onclick=\\\"showWindow(&#39;sendpm&#39;, this.href);\\\" title=\\\"发消息\\\" class=\\\"xi2\\\">发消息</a></li>\\n</ul>\\n</div>\\n</td>\\n<td class=\\\"plc\\\">\\n<div class=\\\"pi\\\">\\n<strong>\\n<a href=\\\"forum.php?mod=redirect&amp;goto=findpost&amp;ptid=1901572&amp;pid=19133954\\\" id=\\\"postnum19133954\\\" onclick=\\\"setCopy(this.href, &#39;帖子地址复制成功&#39;);return false;\\\">\\n沙发</a>\\n</strong>\\n<div class=\\\"pti\\\">\\n<div class=\\\"pdbt\\\">\\n</div>\\n<div class=\\\"authi\\\">\\n<img class=\\\"authicn vm\\\" id=\\\"authicon19133954\\\" src=\\\"static/image/common/online_member.gif\\\"/>\\n<em id=\\\"authorposton19133954\\\">发表于 <span title=\\\"2024-03-02 08:32:16\\\">昨天 08:32</span></em>\\n<span class=\\\"pipe\\\">|</span>\\n<a href=\\\"forum.php?mod=viewthread&amp;tid=1901572&amp;page=1&amp;authorid=446740\\\" rel=\\\"nofollow\\\">只看该作者</a>\\n</div>\\n</div>\\n</div><div class=\\\"pct\\\"><div class=\\\"pcb\\\">\\n<div class=\\\"t_fsz\\\">\\n<table cellspacing=\\\"0\\\" cellpadding=\\\"0\\\"><tbody><tr><td class=\\\"t_f\\\" id=\\\"postmessage_19133954\\\">\\n皮膚很白身材很棒的妹妹 感謝大大分享</td></tr></tbody></table>\\n\\n\\n</div>\\n<div id=\\\"comment_19133954\\\" class=\\\"cm\\\">\\n</div>\\n\\n<div id=\\\"post_rate_div_19133954\\\"></div>\\n</div>\\n</div>\\n\\n</td></tr>"
	xpath := "/tr/td[@class=\"pls\"]"
	//config := "{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\",\"obj\":{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\"}}"
	srt := ""
	resultHandle := "" //`{"value_content":"//title","id_attr":"//title"}`
	fmt.Println(Exe(srt, xpath, html, resultHandle))
}

func TestXpathAndHandle01(t *testing.T) {
	html := `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title>FC2PPV-4322749 初撮影・顔出しけしからん夢袋天然Gcupを抱えた女◯大生！！清楚な顔して色白桃尻 - 亚洲无码原创 -  98堂[原色花堂] -  Powered by Discuz!</title>
<link href="/thread-1901572-1-1.html" rel="canonical" />
<meta name="keywords" content="FC2PPV-4322749 初撮影・顔出しけしからん夢袋天然Gcupを抱えた女◯大生！！清楚な顔して色白桃尻" />
<meta name="description" content="【影片名称】：初撮影・顔出しけしからん夢袋天然Gcupを抱えた女◯大生！！清楚な顔して色白桃尻ボディと狂おしい程プルっンプルンの美巨乳に挟まれ野外パイズリ青姦潮吹き中出し＆追撃中出し！！【出演女优】：无 ... FC2PPV-4322749 初撮影・顔出しけしからん夢袋天然Gcupを抱えた女◯大生！！清楚な顔して色白桃尻 ,98堂[原色花堂]" />
<meta name="generator" content="Discuz! X3.4" />
<meta name="author" content="Discuz! Team and Comsenz UI Team" />
<meta name="copyright" content="2001-2021 Tencent Cloud." />
<meta name="MSSmartTagsPreventParsing" content="True" />
<meta http-equiv="MSThemeCompatible" content="Yes" />
<base href="/" /><link rel="stylesheet" type="text/css" href="data/cache/style_1_common.css?OfY" /><link rel="stylesheet" type="text/css" href="data/cache/style_1_forum_viewthread.css?OfY" /><link rel="stylesheet" id="css_extstyle" type="text/css" href="./template/default/style/t1/style.css" /><script type="text/javascript">var STYLEID = '1', STATICURL = 'static/', IMGDIR = 'static/image/common', VERHASH = 'OfY', charset = 'utf-8', discuz_uid = '0', cookiepre = 'cPNj_2132_', cookiedomain = '', cookiepath = '/', showusercard = '1', attackevasive = '0', disallowfloat = 'newthread', creditnotice = '3|积分|,5|金钱|,7|色币|,8|评分|', defaultstyle = './template/default/style/t1', REPORTURL = 'aHR0cHM6Ly9wY2Zldy5hcHAvdGhyZWFkLTE5MDE1NzItMS0xLmh0bWw=', SITEURL = '/', JSPATH = 'static/js/', CSSPATH = 'data/cache/style_', DYNAMICURL = '';</script>
<script src="static/js/common.js?OfY" type="text/javascript"></script>
<script src="static/libs/fingerprintjs/fp.min.js" type="text/javascript" onload="fingerprint()"></script>
<meta name="application-name" content="98堂[原色花堂]" />
<meta name="msapplication-tooltip" content="98堂[原色花堂]" />
<meta name="msapplication-task" content="name=新帖瀑布流;action-uri=/portal.php;icon-uri=static/image/common/portal.ico" /><meta name="msapplication-task" content="name=论坛;action-uri=/forum.php;icon-uri=static/image/common/bbs.ico" />
<link rel="archives" title="98堂[原色花堂]" href="/archiver/" />
<script src="static/js/forum.js?OfY" type="text/javascript"></script>
</head>

<body id="nv_forum" class="pg_viewthread" onkeydown="if(event.keyCode==27) return false;">
<div id="append_parent"></div><div id="ajaxwaitid"></div>
<div id="toptb" class="cl">
<div class="wp">
<div class="z"><a href="javascript:;"  onclick="setHomepage('/');">设为首页</a><a href="/"  onclick="addFavorite(this.href, '98堂[原色花堂]');return false;">收藏本站</a><script src="static/js/language.js" type="text/javascript"></script>
</div>
<div class="y">
<a id="switchblind" href="javascript:;" onclick="toggleBlind(this)" title="开启辅助访问" class="switchblind"></a>
<a href="javascript:;" id="switchwidth" onclick="widthauto(this)" title="切换到宽版" class="switchwidth">切换到宽版</a>
</div>
</div>
</div>

<div id="qmenu_menu" class="p_pop blk" style="display: none;">
<div class="ptm pbw hm">
请 <a href="javascript:;" class="xi2" onclick="lsSubmit()"><strong>登录</strong></a> 后使用快捷导航<br />没有帐号？<a href="member.php?mod=register" class="xi2 xw1">立即注册</a>
</div>
<div id="fjump_menu" class="btda"></div></div><div id="hd">
<div class="wp">
<div class="hdc cl"><h2><a href="./" title="98堂[原色花堂]"><img src="static/image/common/logo.png" alt="98堂[原色花堂]" border="0" /></a></h2><script src="static/js/logging.js?OfY" type="text/javascript"></script>
<form method="post" autocomplete="off" id="lsform" action="/member.php?mod=logging&amp;action=login&amp;loginsubmit=yes&amp;infloat=yes&amp;lssubmit=yes" onsubmit="return lsSubmit();">
<div class="fastlg cl">
<span id="return_ls" style="display:none"></span>
<div class="y pns">
<table cellspacing="0" cellpadding="0">
<tr>
<td>
<span class="ftid">
<select name="fastloginfield" id="ls_fastloginfield" width="40" tabindex="900">
<option value="username">用户名</option>
<option value="email">Email</option>
</select>
</span>
<script type="text/javascript">simulateSelect('ls_fastloginfield')</script>
</td>
<td><input type="text" name="username" id="ls_username" autocomplete="off" class="px vm" tabindex="901" /></td>
<td class="fastlg_l"><label for="ls_cookietime"><input type="checkbox" name="cookietime" id="ls_cookietime" class="pc" value="2592000" tabindex="903" checked="checked"/>自动登录</label></td>
<td>&nbsp;<a href="javascript:;" onclick="showWindow('login', 'member.php?mod=logging&action=login&viewlostpw=1')">找回密码</a></td>
</tr>
<tr>
<td><label for="ls_password" class="z psw_w">密码</label></td>
<td><input type="password" name="password" id="ls_password" class="px vm" autocomplete="off" tabindex="902" /></td>
<td class="fastlg_l"><button type="submit" class="pn vm" tabindex="904" style="width: 75px;"><em>登录</em></button></td>
<td>&nbsp;<a href="member.php?mod=register" class="xi2 xw1">立即注册</a></td>
</tr>
</table>
<input type="hidden" name="formhash" value="f472fc1f" />
<input type="hidden" name="quickforward" value="yes" />
<input type="hidden" name="handlekey" value="ls" />
</div>
</div>
</form>

</div>

<div id="nv">
<a href="javascript:;" id="qmenu" onmouseover="delayShow(this, function () {showMenu({'ctrlid':'qmenu','pos':'34!','ctrlclass':'a','duration':2});showForummenu(36);})">快捷导航</a>
<ul><li class="a" id="mn_forum" ><a href="forum.php" hidefocus="true" title="BBS"  >论坛<span>BBS</span></a></li><li id="mn_portal" ><a href="portal.php" hidefocus="true" title="Portal"  >新帖瀑布流<span>Portal</span></a></li></ul>
</div>
<div class="p_pop h_pop" id="mn_userapp_menu" style="display: none"></div><div id="mu" class="cl">
</div><div id="scbar" class="cl">
<form id="scbar_form" method="post" autocomplete="off" onsubmit="searchFocus($('scbar_txt'))" action="search.php?searchsubmit=yes" target="_blank">
<input type="hidden" name="mod" id="scbar_mod" value="search" />
<input type="hidden" name="formhash" value="f472fc1f" />
<input type="hidden" name="srchtype" value="title" />
<input type="hidden" name="srhfid" value="36" />
<input type="hidden" name="srhlocality" value="forum::viewthread" />
<table cellspacing="0" cellpadding="0">
<tr>
<td class="scbar_icon_td"></td>
<td class="scbar_txt_td"><input type="text" name="srchtxt" id="scbar_txt" value="请输入搜索内容" autocomplete="off" x-webkit-speech speech /></td>
<td class="scbar_type_td"><a href="javascript:;" id="scbar_type" class="xg1" onclick="showMenu(this.id)" hidefocus="true">搜索</a></td>
<td class="scbar_btn_td"><button type="submit" name="searchsubmit" id="scbar_btn" sc="1" class="pn pnc" value="true"><strong class="xi2">搜索</strong></button></td>
<td class="scbar_hot_td">
<div id="scbar_hot">
</div>
</td>
</tr>
</table>
</form>
</div>
<ul id="scbar_type_menu" class="p_pop" style="display: none;"><li><a href="javascript:;" rel="curforum" fid="36" >本版</a></li><li><a href="javascript:;" rel="forum" class="curtype">帖子</a></li><li><a href="javascript:;" rel="user">用户</a></li></ul>
<script type="text/javascript">
initSearchmenu('scbar', '');
</script>
</div>
</div>


<div id="wp" class="wp">
<div class="h10"></div>
<div class="show-text cl">
        <a href="https://ls399.cc/?shareName=&proxyAccount=99482798&vertical=1#/" target="_blank" class="js-randomBg">拉斯维加斯</a>
        <a href="https://hg9300f.cc:8989/?c=UNMHW" target="_blank" class="js-randomBg">澳门高爆电子</a>
        <a href="https://ciudadpromo.com" target="_blank" class="js-randomBg">亚博美女赌场</a>
        <a href="https://by301270224.cc/?channelCode=by_147" target="_blank" class="js-randomBg">鲍鱼全球黄播</a>
        <a href="https://4qz.me" target="_blank" class="js-randomBg">免费成人乱伦</a>
        <a href="https://572c.tihlrhpe.vip/aff-ucSu" target="_blank" class="js-randomBg">草榴成人免费</a>
        <a href="https://cfslpis.wkvudxj.xyz/sht666-41" target="_blank" class="js-randomBg">免费萝莉禁区</a>
        <a href="https://asfd6ef.com/yj/6475/zw98sh1awjq" target="_blank" class="js-randomBg">免费顶级暗网</a>
        <a href="https://nwhtizm.wcerxcf.xyz/sht888-30" target="_blank" class="js-randomBg"><font color="#00ff00">91免费射区</font></a>
        <a href="https://web.jbwa.ltd" target="_blank" class="js-randomBg">91成人抖音</a>
    </div>
<script type="text/javascript">var fid = parseInt('36'), tid = parseInt('1901572');</script>

<script src="static/js/forum_viewthread.js?OfY" type="text/javascript"></script>
<script type="text/javascript">zoomstatus = parseInt(1);var imagemaxwidth = '600';var aimgcount = new Array();</script>

<style id="diy_style" type="text/css"></style>
<!--[diy=diynavtop]--><div id="diynavtop" class="area"></div><!--[/diy]-->
<div id="pt" class="bm cl">
<div class="z">
<a href="./" class="nvhm" title="首页">98堂[原色花堂]</a><em>&raquo;</em><a href="forum.php">论坛</a> <em>&rsaquo;</em> <a href="forum.php?gid=1">原创BT电影</a> <em>&rsaquo;</em> <a href="forum-36-1.html">亚洲无码原创</a> <em>&rsaquo;</em> <a href="thread-1901572-1-1.html">FC2PPV-4322749 初撮影・顔出しけしからん夢袋天然Gcup ...</a>
</div>
</div>

<style id="diy_style" type="text/css"></style>
<div class="wp">
<!--[diy=diy1]--><div id="diy1" class="area"></div><!--[/diy]-->
</div>

<div id="ct" class="wp cl">
<div id="pgt" class="pgs mbm cl ">
<div class="pgt"><div class="pg"><strong>1</strong><a href="thread-1901572-2-1.html">2</a><a href="thread-1901572-3-1.html">3</a><a href="thread-1901572-4-1.html">4</a><a href="thread-1901572-5-1.html">5</a><a href="thread-1901572-6-1.html">6</a><a href="thread-1901572-7-1.html">7</a><label><input type="text" name="custompage" class="px" size="2" title="输入页码，按回车快速跳转" value="1" onkeydown="if(event.keyCode==13) {window.location='forum.php?mod=viewthread&tid=1901572&amp;extra=page%3D1&amp;page='+this.value;; doane(event);}" /><span title="共 7 页"> / 7 页</span></label><a href="thread-1901572-2-1.html" class="nxt">下一页</a></div></div>
<span class="y pgb" id="visitedforums" onmouseover="$('visitedforums').id = 'visitedforumstmp';this.id = 'visitedforums';showMenu({'ctrlid':this.id,'pos':'34'})"><a href="forum-36-1.html">返回列表</a></span>
<a id="newspecial" onmouseover="$('newspecial').id = 'newspecialtmp';this.id = 'newspecial';showMenu({'ctrlid':this.id})" onclick="showWindow('newthread', 'forum.php?mod=post&action=newthread&fid=36')" href="javascript:;" title="发新帖"><img src="static/image/common/pn_post.png" alt="发新帖" /></a></div>



<div id="postlist" class="pl bm">
<table cellspacing="0" cellpadding="0">
<tr>
<td class="pls ptn pbn">
<div class="hm ptn">
<span class="xg1">查看:</span> <span class="xi1">32442</span><span class="pipe">|</span><span class="xg1">回复:</span> <span class="xi1">66</span>
</div>
</td>
<td class="plc ptm pbn vwthd">
<div class="y">
<a href="forum.php?mod=viewthread&amp;action=printable&amp;tid=1901572" title="打印" target="_blank"><img src="static/image/common/print.png" alt="打印" class="vm" /></a>
<a href="forum.php?mod=redirect&amp;goto=nextoldset&amp;tid=1901572" title="上一主题"><img src="static/image/common/thread-prev.png" alt="上一主题" class="vm" /></a>
<a href="forum.php?mod=redirect&amp;goto=nextnewset&amp;tid=1901572" title="下一主题"><img src="static/image/common/thread-next.png" alt="下一主题" class="vm" /></a>
</div>
<h1 class="ts">
<a href="forum.php?mod=forumdisplay&amp;fid=36&amp;filter=typeid&amp;typeid=368">[FC2PPV]</a>
<span id="thread_subject">FC2PPV-4322749 初撮影・顔出しけしからん夢袋天然Gcupを抱えた女◯大生！！清楚な顔して色白桃尻</span>
</h1>
<span class="xg1">
&nbsp;<img src="static/image/common/hot_2.gif" alt="" title="热度: 102" />
<a href="thread-1901572-1-1.html" onclick="return copyThreadUrl(this, '98堂[原色花堂]')" >[复制链接]</a>
</span>
</td>
</tr>
</table>


<table cellspacing="0" cellpadding="0" class="ad">
<tr>
<td class="pls">
</td>
<td class="plc">
</td>
</tr>
</table><style>
.show-text4{position:relative;padding-left:170px;min-height:52px}
.show-text4 .show-info{position:absolute;top:0;left:0;width:160px;height:100%;text-align:center;display:block;border-right:1px solid #c2d5e3}
.show-text4 .show-info .avatar{position:absolute;top:50%;left:50%;margin-top:-23px;margin-left:-23px;width:46px;height:46px;line-height:46px;-webkit-border-radius:5rem;border-radius:5rem;background:#c1c1c1;color:#fff;text-align:center;font-size:18px}
.show-text4 .show-main{padding:5px 10px}
.show-text4 .item{width:50%;float:left;display:block;text-decoration:none}
.show-text4 .item h4{font-size:16px;color:#36f}
.show-text4 .item p{color:#777}
</style><div id="post_19133577" ><table id="pid19133577" class="plhin" summary="pid19133577" cellspacing="0" cellpadding="0">
<tr>
<a name="newpost"></a> <td class="pls" rowspan="2">
<div id="favatar19133577" class="pls cl favatar">
<div class="pi">
<div class="authi"><a href="space-uid-420458.html" target="_blank" class="xw1" style="color: #0033FF">junmin1995</a>
</div>
</div>
<div class="p_pop blk bui card_gender_0" id="userinfo19133577" style="display: none; margin-top: -11px;">
<div class="m z">
<div id="userinfo19133577_ma"></div>
</div>
<div class="i y">
<div>
<strong><a href="space-uid-420458.html" target="_blank" class="xi2" style="color: #0033FF">junmin1995</a></strong>
<em>当前在线</em>
</div><dl class="cl">
<dt>积分</dt><dd><a href="home.php?mod=space&uid=420458&do=profile" target="_blank" class="xi2">118336</a></dd>
</dl><div class="imicn">
<a href="home.php?mod=space&amp;uid=420458&amp;do=profile" target="_blank" title="查看详细资料"><img src="static/image/common/userinfo.gif" alt="查看详细资料" /></a>
</div>
<div id="avatarfeed"><span id="threadsortswait"></span></div>
</div>
</div>
<div>
<div class="avatar" onmouseover="showauthor(this, 'userinfo19133577')"><a href="space-uid-420458.html" class="avtm" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/42/04/58_avatar_middle.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_middle.gif'" /></a></div>
</div>
<div class="tns xg2"><table cellspacing="0" cellpadding="0"><th><p><a href="home.php?mod=space&uid=420458&do=profile" class="xi2"><span title="119151">11万</span></a></p>金钱</th><th><p><a href="home.php?mod=space&uid=420458&do=profile" class="xi2">0</a></p>色币</th><td><p><a href="home.php?mod=space&uid=420458&do=profile" class="xi2"><span title="653892">65万</span></a></p>评分</td></table></div>

<p><em><a href="home.php?mod=spacecp&amp;ac=usergroup&amp;gid=3" target="_blank"><font color="#0033FF">版主</font></a></em></p>


<dl class="pil cl">
  <dt>积分</dt><dd><a href="home.php?mod=space&uid=420458&do=profile" target="_blank" class="xi2">118336</a></dd>
</dl>


<ul class="xl xl2 o cl">
<li class="pm2"><a href="home.php?mod=spacecp&amp;ac=pm&amp;op=showmsg&amp;handlekey=showmsg_420458&amp;touid=420458&amp;pmid=0&amp;daterange=2&amp;pid=19133577&amp;tid=1901572" onclick="showWindow('sendpm', this.href);" title="发消息" class="xi2">发消息</a></li>
</ul>
</div>
</td>
<td class="plc">
<div class="pi">
<div id="fj" class="y">
<label class="z">电梯直达</label>
<input type="text" class="px p_fre z" size="2" onkeyup="$('fj_btn').href='forum.php?mod=redirect&ptid=1901572&authorid=0&postno='+this.value" onkeydown="if(event.keyCode==13) {window.location=$('fj_btn').href;return false;}" title="跳转到指定楼层" />
<a href="javascript:;" id="fj_btn" class="z" title="跳转到指定楼层"><img src="static/image/common/fj_btn.png" alt="跳转到指定楼层" class="vm" /></a>
</div>
<strong>
<a href="thread-1901572-1-1.html"   id="postnum19133577" onclick="setCopy(this.href, '帖子地址复制成功');return false;">
楼主</a>
</strong>
<div class="pti">
<div class="pdbt">
</div>
<div class="authi">
<img class="authicn vm" id="authicon19133577" src="static/image/common/online_moderator.gif" />
<em id="authorposton19133577">发表于 <span title="2024-03-02 08:18:40">昨天&nbsp;08:18</span></em>
<span class="pipe">|</span>
<a href="forum.php?mod=viewthread&amp;tid=1901572&amp;page=1&amp;authorid=420458" rel="nofollow">只看该作者</a>
<span class="pipe">|</span><a href="forum.php?mod=viewthread&amp;tid=1901572&amp;from=album">只看大图</a>
<span class="none"><img src="static/image/common/arw_r.gif" class="vm" alt="回帖奖励" /></span>
<span class="pipe show">|</span><a href="forum.php?mod=viewthread&amp;tid=1901572&amp;extra=page%3D1&amp;ordertype=1"  class="show">倒序浏览</a>
<span class="pipe show">|</span><a href="javascript:;" onclick="readmode($('thread_subject').innerHTML, 19133577);" class="show">阅读模式</a>
</div>
</div>
</div><div class="pct"><style type="text/css">.pcb{margin-right:0}</style><div class="pcb">
<div class="t_fsz">
<table cellspacing="0" cellpadding="0"><tr><td class="t_f" id="postmessage_19133577">
【影片名称】：初撮影・顔出しけしからん夢袋天然Gcupを抱えた女◯大生！！清楚な顔して色白桃尻ボディと狂おしい程プルっンプルンの美巨乳に挟まれ野外パイズリ青姦潮吹き中出し＆追撃中出し！！<br />
【出演女优】：无名<br />
【影片容量】：3.76G<br />
【是否有码】：无码<br />
【种子期限】：5种或健康度1000<br />
【下载工具】：比特彗星 比特精灵 uTorrent QBittorrent 迅雷极速版 闪电下载【请不要用迅雷官方版下载，官方版本已经被屏蔽】<br />
【影片预览】：看不到图请挂代理或点右键显示图片<br />

<ignore_js_op>

<img aid="8191095" src="static/image/common/none.gif" zoomfile="https://1szbg.app/tupian/forum/202403/02/081821d7ppttvtlpp8hd8h.jpg" file="https://1szbg.app/tupian/forum/202403/02/081821d7ppttvtlpp8hd8h.jpg" class="zoom" onclick="zoom(this, this.src, 0, 0, 0)" width="600" id="aimg_8191095" inpost="1" onmouseover="showMenu({'ctrlid':this.id,'pos':'12'})" />

<div class="tip tip_4 aimg_tip" id="aimg_8191095_menu" style="position: absolute; display: none" disautofocus="true">
<div class="xs0">
<p><strong>FC2PPV-4322749.jpg</strong> <em class="xg1">(320.9 KB, 下载次数: 0)</em></p>
<p>

<a href="https://1szbg.app/tupian/forum/202403/02/081821d7ppttvtlpp8hd8h.jpg" target="_blank">下载附件</a>

</p>

<p class="xg1 y"><span title="2024-03-02 08:18">昨天&nbsp;08:18</span> 上传</p>

</div>
<div class="tip_horn"></div>
</div>

</ignore_js_op>
<br />

<ignore_js_op>

<img aid="8191096" src="static/image/common/none.gif" zoomfile="https://1szbg.app/tupian/forum/202403/02/081821ufwpfbwefxpapu7e.jpg" file="https://1szbg.app/tupian/forum/202403/02/081821ufwpfbwefxpapu7e.jpg" class="zoom" onclick="zoom(this, this.src, 0, 0, 0)" width="600" id="aimg_8191096" inpost="1" onmouseover="showMenu({'ctrlid':this.id,'pos':'12'})" />

<div class="tip tip_4 aimg_tip" id="aimg_8191096_menu" style="position: absolute; display: none" disautofocus="true">
<div class="xs0">
<p><strong>FC2PPV-4322749.mp4.jpg</strong> <em class="xg1">(179.3 KB, 下载次数: 0)</em></p>
<p>

<a href="https://1szbg.app/tupian/forum/202403/02/081821ufwpfbwefxpapu7e.jpg" target="_blank">下载附件</a>

</p>

<p class="xg1 y"><span title="2024-03-02 08:18">昨天&nbsp;08:18</span> 上传</p>

</div>
<div class="tip_horn"></div>
</div>

</ignore_js_op>
<br />
磁力链接：<br />
<div class="blockcode"><div id="code_iML"><ol><li>magnet:?xt=urn:btih:B0E49A7D8AA7551EFEC9AE475ACC7006E32F949F</ol></div><em onclick="copycode($('code_iML'));">复制代码</em></div><br />

<br />

<br />
</td></tr></table>

<div class="modact"><a href="forum.php?mod=misc&amp;action=viewthreadmod&amp;tid=1901572" title="帖子模式" onclick="showWindow('viewthreadmod', this.href)">本主题由 junmin1995 于 <span title="2024-03-02 08:21">昨天&nbsp;08:21</span> 限时高亮</a></div><div class="pattl">

<ignore_js_op>
<dl class="tattl">
<dt>
<img src="static/image/filetype/torrent.gif" border="0" class="vm" alt="" />
</dt>
<dd>
<p class="attnm">

<a href="forum.php?mod=attachment&aid=ODE5MTA5N3w3ZTg0MDI4NnwxNzA5NDY4OTU2fDB8MTkwMTU3Mg%3D%3D" onmouseover="showMenu({'ctrlid':this.id,'pos':'12'})" id="aid8191097" target="_blank">FC2PPV-4322749.torrent</a>

<div class="tip tip_4" id="aid8191097_menu" style="display: none" disautofocus="true">
<div class="tip_c">
<p class="y"><span title="2024-03-02 08:18">昨天&nbsp;08:18</span> 上传</p>
<p>点击文件名下载附件</p>

</div>
<div class="tip_horn"></div>
</div>
</p>
<p>21.03 KB, 下载次数: 2599</p>
<p>

</p>


</dd>
</dl>
</ignore_js_op>
</div>

</div>
<div id="comment_19133577" class="cm">
</div>

<h3 class="psth xs1"><span class="icon_ring vm"></span>评分</h3>
<dl id="ratelog_19133577" class="rate">
<dd style="margin:0">
<div id="post_rate_19133577"></div>
<table class="ratl">
<tr>
<th class="xw1" width="120"><a href="forum.php?mod=misc&amp;action=viewratings&amp;tid=1901572&amp;pid=19133577" onclick="showWindow('viewratings', this.href)" title="查看全部评分"> 参与人数 <span class="xi1">12</span></a></th><th class="xw1" width="80">评分 <i><span class="xi1">+28</span></i></th>
<th>
<a href="javascript:;" onclick="toggleRatelogCollapse('ratelog_19133577', this);" class="y xi2 op">收起</a>
<i class="txt_h">理由</i>
</th>
</tr>
<tbody class="ratl_l"><tr id="rate_19133577_399017">
<td>
<a href="space-uid-399017.html" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/39/90/17_avatar_small.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_small.gif'" /></a> <a href="space-uid-399017.html" target="_blank">xx86376150</a>
</td><td class="xi1"> + 2</td>
<td class="xg1">很给力!</td>
</tr>
<tr id="rate_19133577_483682">
<td>
<a href="space-uid-483682.html" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/48/36/82_avatar_small.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_small.gif'" /></a> <a href="space-uid-483682.html" target="_blank">12345tao</a>
</td><td class="xi1"> + 1</td>
<td class="xg1">很给力!</td>
</tr>
<tr id="rate_19133577_439033">
<td>
<a href="space-uid-439033.html" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/43/90/33_avatar_small.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_small.gif'" /></a> <a href="space-uid-439033.html" target="_blank">lsp1221</a>
</td><td class="xi1"> + 2</td>
<td class="xg1"></td>
</tr>
<tr id="rate_19133577_477287">
<td>
<a href="space-uid-477287.html" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/47/72/87_avatar_small.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_small.gif'" /></a> <a href="space-uid-477287.html" target="_blank">kkh1234567</a>
</td><td class="xi1"> + 4</td>
<td class="xg1"></td>
</tr>
<tr id="rate_19133577_448316">
<td>
<a href="space-uid-448316.html" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/44/83/16_avatar_small.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_small.gif'" /></a> <a href="space-uid-448316.html" target="_blank">dhbrhd</a>
</td><td class="xi1"> + 3</td>
<td class="xg1">很给力!</td>
</tr>
<tr id="rate_19133577_424388">
<td>
<a href="space-uid-424388.html" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/42/43/88_avatar_small.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_small.gif'" /></a> <a href="space-uid-424388.html" target="_blank">huamanlou010</a>
</td><td class="xi1"> + 3</td>
<td class="xg1">很给力!</td>
</tr>
<tr id="rate_19133577_444566">
<td>
<a href="space-uid-444566.html" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/44/45/66_avatar_small.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_small.gif'" /></a> <a href="space-uid-444566.html" target="_blank">无可言喻</a>
</td><td class="xi1"> + 2</td>
<td class="xg1">很给力!</td>
</tr>
<tr id="rate_19133577_483556">
<td>
<a href="space-uid-483556.html" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/48/35/56_avatar_small.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_small.gif'" /></a> <a href="space-uid-483556.html" target="_blank">foolday</a>
</td><td class="xi1"> + 3</td>
<td class="xg1"></td>
</tr>
<tr id="rate_19133577_482230">
<td>
<a href="space-uid-482230.html" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/48/22/30_avatar_small.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_small.gif'" /></a> <a href="space-uid-482230.html" target="_blank">yihongyuan010</a>
</td><td class="xi1"> + 3</td>
<td class="xg1">很给力!</td>
</tr>
<tr id="rate_19133577_273594">
<td>
<a href="space-uid-273594.html" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/27/35/94_avatar_small.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_small.gif'" /></a> <a href="space-uid-273594.html" target="_blank">a1a1_a1</a>
</td><td class="xi1"> + 1</td>
<td class="xg1"></td>
</tr>
<tr id="rate_19133577_483005">
<td>
<a href="space-uid-483005.html" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/48/30/05_avatar_small.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_small.gif'" /></a> <a href="space-uid-483005.html" target="_blank">Dickdark</a>
</td><td class="xi1"> + 1</td>
<td class="xg1">很给力!</td>
</tr>
<tr id="rate_19133577_446740">
<td>
<a href="space-uid-446740.html" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/44/67/40_avatar_small.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_small.gif'" /></a> <a href="space-uid-446740.html" target="_blank">blue2023</a>
</td><td class="xi1"> + 3</td>
<td class="xg1"></td>
</tr>
</tbody>
</table>
<p class="ratc">
<a href="forum.php?mod=misc&amp;action=viewratings&amp;tid=1901572&amp;pid=19133577" onclick="showWindow('viewratings', this.href)" title="查看全部评分" class="xi2">查看全部评分</a>
</p>
</dd>
</dl>
</div>
</div>

</td></tr>
<tr><td class="plc plm">

<div id="p_btn" class="mtw mbm hm cl">

<a href="home.php?mod=spacecp&amp;ac=favorite&amp;type=thread&amp;id=1901572&amp;formhash=f472fc1f" id="k_favorite" onclick="showWindow(this.id, this.href, 'get', 0);" onmouseover="this.title = $('favoritenumber').innerHTML + ' 人收藏'" title="收藏本帖"><i><img src="static/image/common/fav.gif" alt="收藏" />收藏<span id="favoritenumber">33</span></i></a>
</div>

</td>
</tr>
<tr id="_postposition19133577"></tr>
<tr>
<td class="pls"></td>
<td class="plc" style="overflow:visible;">
<div class="po hin">
<div class="pob cl">
<em>
<a class="fastre" href="forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1901572&amp;reppost=19133577&amp;extra=page%3D1&amp;page=1" onclick="showWindow('reply', this.href)">回复</a>
</em>

<p>
<a href="javascript:;" id="mgc_post_19133577" onmouseover="showMenu(this.id)" class="showmenu">使用道具</a>
<a href="javascript:;" onclick="showWindow('miscreport19133577', 'misc.php?mod=report&rtype=post&rid=19133577&tid=1901572&fid=36', 'get', -1);return false;">举报</a>
</p>

<ul id="mgc_post_19133577_menu" class="p_pop mgcmn" style="display: none;">
</ul>
<script type="text/javascript" reload="1">checkmgcmn('post_19133577')</script>
</div>
</div>
</td>
</tr>
<tr class="ad">
<td class="pls">
</td>
<td class="plc">
</td>
</tr>
</table>
<script type="text/javascript" reload="1">
aimgcount[19133577] = ['8191095','8191096'];
attachimggroup(19133577);
var aimgfid = 0;
</script>
</div>

<div class="show-text2 pad-tb-10"><a href="https://ls399.cc/?shareName=&proxyAccount=99482798&vertical=1#/" target="_blank">拉斯维加斯</a>
<a href="https://hg9300f.cc:8989/?c=UNMHW" target="_blank">澳门高爆电子</a>
<a href="https://ciudadpromo.com" target="_blank">亚博美女赌场</a>
<a href="https://by301270224.cc/?channelCode=by_147" target="_blank">鲍鱼全球黄播</a>
<a href="https://4qz.me" target="_blank">免费成人乱伦</a>
<a href="https://572c.tihlrhpe.vip/aff-ucSu" target="_blank">草榴成人免费</a>
<a href="https://cfslpis.wkvudxj.xyz/sht666-41" target="_blank">免费萝莉禁区</a>
<a href="https://asfd6ef.com/yj/6475/zw98sh1awjq" target="_blank">免费顶级暗网</a>
<a href="https://nwhtizm.wcerxcf.xyz/sht888-30" target="_blank">91免费射区</a>
<a href="https://web.jbwa.ltd" target="_blank">91成人抖音</a>
</div>
<table class="plhin">
<tr class="ad">
<td class="pls"></td>
<td class="plc"></td>
</tr>
</table>
<div class="show-text4 pad-tb-10">
<div class="show-info">
<div class="avatar">AD</div>
</div>
<div class="show-main"><div><a href="https://27.25.129.245:3306/#/?uid=50652642" class="item js-appJump">
    <h4>鲍鱼盒子，全网高清直播资源‖自拍视频在线看</h4>
    <p>注册免费观看，邀请好友送现金红包。</p>
</a>
<a href="https://27.25.129.245:3306/#/?uid=50652642" class="item js-appJump">
    <h4>鲍鱼盒子，全网高清直播资源‖自拍视频在线看</h4>
    <p>注册免费观看，邀请好友送现金红包。</p>
</a></div></div>
</div>
<table class="plhin">
<tr class="ad">
<td class="pls"></td>
<td class="plc"></td>
</tr>
</table>
<div id="post_19133954" ><table id="pid19133954" class="plhin" summary="pid19133954" cellspacing="0" cellpadding="0">
<tr>
 <td class="pls" rowspan="2">
<div id="favatar19133954" class="pls cl favatar">
<div class="pi">
<div class="authi"><a href="space-uid-446740.html" target="_blank" class="xw1">blue2023</a>
</div>
</div>
<div class="p_pop blk bui card_gender_0" id="userinfo19133954" style="display: none; margin-top: -11px;">
<div class="m z">
<div id="userinfo19133954_ma"></div>
</div>
<div class="i y">
<div>
<strong><a href="space-uid-446740.html" target="_blank" class="xi2">blue2023</a></strong>
<em>当前离线</em>
</div><dl class="cl">
<dt>积分</dt><dd><a href="home.php?mod=space&uid=446740&do=profile" target="_blank" class="xi2">259</a></dd>
</dl><div class="imicn">
<a href="home.php?mod=space&amp;uid=446740&amp;do=profile" target="_blank" title="查看详细资料"><img src="static/image/common/userinfo.gif" alt="查看详细资料" /></a>
</div>
<div id="avatarfeed"><span id="threadsortswait"></span></div>
</div>
</div>
<div>
<div class="avatar" onmouseover="showauthor(this, 'userinfo19133954')"><a href="space-uid-446740.html" class="avtm" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/44/67/40_avatar_middle.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_middle.gif'" /></a></div>
</div>
<div class="tns xg2"><table cellspacing="0" cellpadding="0"><th><p><a href="home.php?mod=space&uid=446740&do=profile" class="xi2">724</a></p>金钱</th><th><p><a href="home.php?mod=space&uid=446740&do=profile" class="xi2">0</a></p>色币</th><td><p><a href="home.php?mod=space&uid=446740&do=profile" class="xi2">9</a></p>评分</td></table></div>

<p><em><a href="home.php?mod=spacecp&amp;ac=usergroup&amp;gid=15" target="_blank">Lv6 后起之秀</a></em></p>


<p><span class="pbg2"  id="upgradeprogress_19133954" onmouseover="showMenu({'ctrlid':this.id, 'pos':'12!', 'menuid':'g_up19133954_menu'});"><span class="pbr2" style="width:19%;"></span></span></p>
<div id="g_up19133954_menu" class="tip tip_4" style="display: none;"><div class="tip_horn"></div><div class="tip_c">Lv6 后起之秀, 积分 259, 距离下一级还需 241 积分</div></div>

<dl class="pil cl">
  <dt>积分</dt><dd><a href="home.php?mod=space&uid=446740&do=profile" target="_blank" class="xi2">259</a></dd>
</dl>


<ul class="xl xl2 o cl">
<li class="pm2"><a href="home.php?mod=spacecp&amp;ac=pm&amp;op=showmsg&amp;handlekey=showmsg_446740&amp;touid=446740&amp;pmid=0&amp;daterange=2&amp;pid=19133954&amp;tid=1901572" onclick="showWindow('sendpm', this.href);" title="发消息" class="xi2">发消息</a></li>
</ul>
</div>
</td>
<td class="plc">
<div class="pi">
<strong>
<a href="forum.php?mod=redirect&goto=findpost&ptid=1901572&pid=19133954"   id="postnum19133954" onclick="setCopy(this.href, '帖子地址复制成功');return false;">
沙发</a>
</strong>
<div class="pti">
<div class="pdbt">
</div>
<div class="authi">
<img class="authicn vm" id="authicon19133954" src="static/image/common/online_member.gif" />
<em id="authorposton19133954">发表于 <span title="2024-03-02 08:32:16">昨天&nbsp;08:32</span></em>
<span class="pipe">|</span>
<a href="forum.php?mod=viewthread&amp;tid=1901572&amp;page=1&amp;authorid=446740" rel="nofollow">只看该作者</a>
</div>
</div>
</div><div class="pct"><div class="pcb">
<div class="t_fsz">
<table cellspacing="0" cellpadding="0"><tr><td class="t_f" id="postmessage_19133954">
皮膚很白身材很棒的妹妹 感謝大大分享</td></tr></table>


</div>
<div id="comment_19133954" class="cm">
</div>

<div id="post_rate_div_19133954"></div>
</div>
</div>

</td></tr>
<tr><td class="plc plm">

<div id="p_btn" class="mtw mbm hm cl">
</div>

</td>
</tr>
<tr id="_postposition19133954"></tr>
<tr>
<td class="pls"></td>
<td class="plc" style="overflow:visible;">
<div class="po hin">
<div class="pob cl">
<em>
<a class="fastre" href="forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1901572&amp;repquote=19133954&amp;extra=page%3D1&amp;page=1" onclick="showWindow('reply', this.href)">回复</a>
</em>

<p>
<a href="javascript:;" id="mgc_post_19133954" onmouseover="showMenu(this.id)" class="showmenu">使用道具</a>
<a href="javascript:;" onclick="showWindow('miscreport19133954', 'misc.php?mod=report&rtype=post&rid=19133954&tid=1901572&fid=36', 'get', -1);return false;">举报</a>
</p>

<ul id="mgc_post_19133954_menu" class="p_pop mgcmn" style="display: none;">
</ul>
<script type="text/javascript" reload="1">checkmgcmn('post_19133954')</script>
</div>
</div>
</td>
</tr>
<tr class="ad">
<td class="pls">
</td>
<td class="plc">
</td>
</tr>
</table>
</div>

<div id="post_19134901" ><table id="pid19134901" class="plhin" summary="pid19134901" cellspacing="0" cellpadding="0">
<tr>
 <td class="pls" rowspan="2">
<div id="favatar19134901" class="pls cl favatar">
<div class="pi">
<div class="authi"><a href="space-uid-483005.html" target="_blank" class="xw1">Dickdark</a>
</div>
</div>
<div class="p_pop blk bui card_gender_0" id="userinfo19134901" style="display: none; margin-top: -11px;">
<div class="m z">
<div id="userinfo19134901_ma"></div>
</div>
<div class="i y">
<div>
<strong><a href="space-uid-483005.html" target="_blank" class="xi2">Dickdark</a></strong>
<em>当前离线</em>
</div><dl class="cl">
<dt>积分</dt><dd><a href="home.php?mod=space&uid=483005&do=profile" target="_blank" class="xi2">43</a></dd>
</dl><div class="imicn">
<a href="home.php?mod=space&amp;uid=483005&amp;do=profile" target="_blank" title="查看详细资料"><img src="static/image/common/userinfo.gif" alt="查看详细资料" /></a>
</div>
<div id="avatarfeed"><span id="threadsortswait"></span></div>
</div>
</div>
<div>
<div class="avatar" onmouseover="showauthor(this, 'userinfo19134901')"><a href="space-uid-483005.html" class="avtm" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/48/30/05_avatar_middle.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_middle.gif'" /></a></div>
</div>
<div class="tns xg2"><table cellspacing="0" cellpadding="0"><th><p><a href="home.php?mod=space&uid=483005&do=profile" class="xi2">36</a></p>金钱</th><th><p><a href="home.php?mod=space&uid=483005&do=profile" class="xi2">0</a></p>色币</th><td><p><a href="home.php?mod=space&uid=483005&do=profile" class="xi2">0</a></p>评分</td></table></div>

<p><em><a href="home.php?mod=spacecp&amp;ac=usergroup&amp;gid=12" target="_blank">Lv3 江湖小虾</a></em></p>


<p><span class="pbg2"  id="upgradeprogress_19134901" onmouseover="showMenu({'ctrlid':this.id, 'pos':'12!', 'menuid':'g_up19134901_menu'});"><span class="pbr2" style="width:65%;"></span></span></p>
<div id="g_up19134901_menu" class="tip tip_4" style="display: none;"><div class="tip_horn"></div><div class="tip_c">Lv3 江湖小虾, 积分 43, 距离下一级还需 7 积分</div></div>

<dl class="pil cl">
  <dt>积分</dt><dd><a href="home.php?mod=space&uid=483005&do=profile" target="_blank" class="xi2">43</a></dd>
</dl>


<ul class="xl xl2 o cl">
<li class="pm2"><a href="home.php?mod=spacecp&amp;ac=pm&amp;op=showmsg&amp;handlekey=showmsg_483005&amp;touid=483005&amp;pmid=0&amp;daterange=2&amp;pid=19134901&amp;tid=1901572" onclick="showWindow('sendpm', this.href);" title="发消息" class="xi2">发消息</a></li>
</ul>
</div>
</td>
<td class="plc">
<div class="pi">
<strong>
<a href="forum.php?mod=redirect&goto=findpost&ptid=1901572&pid=19134901"   id="postnum19134901" onclick="setCopy(this.href, '帖子地址复制成功');return false;">
板凳</a>
</strong>
<div class="pti">
<div class="pdbt">
</div>
<div class="authi">
<img class="authicn vm" id="authicon19134901" src="static/image/common/online_member.gif" />
<em id="authorposton19134901">发表于 <span title="2024-03-02 09:04:21">昨天&nbsp;09:04</span></em>
<span class="xg1">来自手机</span>
<span class="pipe">|</span>
<a href="forum.php?mod=viewthread&amp;tid=1901572&amp;page=1&amp;authorid=483005" rel="nofollow">只看该作者</a>
</div>
</div>
</div><div class="pct"><div class="pcb">
<div class="t_fsz">
<table cellspacing="0" cellpadding="0"><tr><td class="t_f" id="postmessage_19134901">
今日最佳</td></tr></table>


</div>
<div id="comment_19134901" class="cm">
</div>

<div id="post_rate_div_19134901"></div>
</div>
</div>

</td></tr>
<tr><td class="plc plm">

<div id="p_btn" class="mtw mbm hm cl">
</div>

</td>
</tr>
<tr id="_postposition19134901"></tr>
<tr>
<td class="pls"></td>
<td class="plc" style="overflow:visible;">
<div class="po hin">
<div class="pob cl">
<em>
<a class="fastre" href="forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1901572&amp;repquote=19134901&amp;extra=page%3D1&amp;page=1" onclick="showWindow('reply', this.href)">回复</a>
</em>

<p>
<a href="javascript:;" id="mgc_post_19134901" onmouseover="showMenu(this.id)" class="showmenu">使用道具</a>
<a href="javascript:;" onclick="showWindow('miscreport19134901', 'misc.php?mod=report&rtype=post&rid=19134901&tid=1901572&fid=36', 'get', -1);return false;">举报</a>
</p>

<ul id="mgc_post_19134901_menu" class="p_pop mgcmn" style="display: none;">
</ul>
<script type="text/javascript" reload="1">checkmgcmn('post_19134901')</script>
</div>
</div>
</td>
</tr>
<tr class="ad">
<td class="pls">
</td>
<td class="plc">
</td>
</tr>
</table>
</div>

<div class="show-text4 pad-tb-10">
<div class="show-info">
<div class="avatar">AD</div>
</div>
<div class="show-main"><div><a href="https://27.25.129.245:3306/#/?uid=50652642" class="item js-appJump">
    <h4>鲍鱼盒子，全网高清直播资源‖自拍视频在线看</h4>
    <p>注册免费观看，邀请好友送现金红包。</p>
</a>
<a href="https://27.25.129.245:3306/#/?uid=50652642" class="item js-appJump">
    <h4>鲍鱼盒子，全网高清直播资源‖自拍视频在线看</h4>
    <p>注册免费观看，邀请好友送现金红包。</p>
</a></div></div>
</div>
<table class="plhin">
<tr class="ad">
<td class="pls"></td>
<td class="plc"></td>
</tr>
</table>
<div id="post_19135338" ><table id="pid19135338" class="plhin" summary="pid19135338" cellspacing="0" cellpadding="0">
<tr>
 <td class="pls" rowspan="2">
<div id="favatar19135338" class="pls cl favatar">
<div class="pi">
<div class="authi"><a href="space-uid-395648.html" target="_blank" class="xw1">zsfq233</a>
</div>
</div>
<div class="p_pop blk bui card_gender_0" id="userinfo19135338" style="display: none; margin-top: -11px;">
<div class="m z">
<div id="userinfo19135338_ma"></div>
</div>
<div class="i y">
<div>
<strong><a href="space-uid-395648.html" target="_blank" class="xi2">zsfq233</a></strong>
<em>当前离线</em>
</div><dl class="cl">
<dt>积分</dt><dd><a href="home.php?mod=space&uid=395648&do=profile" target="_blank" class="xi2">127</a></dd>
</dl><div class="imicn">
<a href="home.php?mod=space&amp;uid=395648&amp;do=profile" target="_blank" title="查看详细资料"><img src="static/image/common/userinfo.gif" alt="查看详细资料" /></a>
</div>
<div id="avatarfeed"><span id="threadsortswait"></span></div>
</div>
</div>
<div>
<div class="avatar" onmouseover="showauthor(this, 'userinfo19135338')"><a href="space-uid-395648.html" class="avtm" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/39/56/48_avatar_middle.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_middle.gif'" /></a></div>
</div>
<div class="tns xg2"><table cellspacing="0" cellpadding="0"><th><p><a href="home.php?mod=space&uid=395648&do=profile" class="xi2">3219</a></p>金钱</th><th><p><a href="home.php?mod=space&uid=395648&do=profile" class="xi2">0</a></p>色币</th><td><p><a href="home.php?mod=space&uid=395648&do=profile" class="xi2">2</a></p>评分</td></table></div>

<p><em><a href="home.php?mod=spacecp&amp;ac=usergroup&amp;gid=14" target="_blank">Lv5 小有名气</a></em></p>


<p><span class="pbg2"  id="upgradeprogress_19135338" onmouseover="showMenu({'ctrlid':this.id, 'pos':'12!', 'menuid':'g_up19135338_menu'});"><span class="pbr2" style="width:27%;"></span></span></p>
<div id="g_up19135338_menu" class="tip tip_4" style="display: none;"><div class="tip_horn"></div><div class="tip_c">Lv5 小有名气, 积分 127, 距离下一级还需 73 积分</div></div>

<dl class="pil cl">
  <dt>积分</dt><dd><a href="home.php?mod=space&uid=395648&do=profile" target="_blank" class="xi2">127</a></dd>
</dl>


<ul class="xl xl2 o cl">
<li class="pm2"><a href="home.php?mod=spacecp&amp;ac=pm&amp;op=showmsg&amp;handlekey=showmsg_395648&amp;touid=395648&amp;pmid=0&amp;daterange=2&amp;pid=19135338&amp;tid=1901572" onclick="showWindow('sendpm', this.href);" title="发消息" class="xi2">发消息</a></li>
</ul>
</div>
</td>
<td class="plc">
<div class="pi">
<strong>
<a href="forum.php?mod=redirect&goto=findpost&ptid=1901572&pid=19135338"   id="postnum19135338" onclick="setCopy(this.href, '帖子地址复制成功');return false;">
地板</a>
</strong>
<div class="pti">
<div class="pdbt">
</div>
<div class="authi">
<img class="authicn vm" id="authicon19135338" src="static/image/common/online_member.gif" />
<em id="authorposton19135338">发表于 <span title="2024-03-02 09:20:13">昨天&nbsp;09:20</span></em>
<span class="xg1">来自手机</span>
<span class="pipe">|</span>
<a href="forum.php?mod=viewthread&amp;tid=1901572&amp;page=1&amp;authorid=395648" rel="nofollow">只看该作者</a>
</div>
</div>
</div><div class="pct"><div class="pcb">
<div class="t_fsz">
<table cellspacing="0" cellpadding="0"><tr><td class="t_f" id="postmessage_19135338">
感谢分享</td></tr></table>


</div>
<div id="comment_19135338" class="cm">
</div>

<div id="post_rate_div_19135338"></div>
</div>
</div>

</td></tr>
<tr><td class="plc plm">

<div id="p_btn" class="mtw mbm hm cl">
</div>

</td>
</tr>
<tr id="_postposition19135338"></tr>
<tr>
<td class="pls"></td>
<td class="plc" style="overflow:visible;">
<div class="po hin">
<div class="pob cl">
<em>
<a class="fastre" href="forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1901572&amp;repquote=19135338&amp;extra=page%3D1&amp;page=1" onclick="showWindow('reply', this.href)">回复</a>
</em>

<p>
<a href="javascript:;" id="mgc_post_19135338" onmouseover="showMenu(this.id)" class="showmenu">使用道具</a>
<a href="javascript:;" onclick="showWindow('miscreport19135338', 'misc.php?mod=report&rtype=post&rid=19135338&tid=1901572&fid=36', 'get', -1);return false;">举报</a>
</p>

<ul id="mgc_post_19135338_menu" class="p_pop mgcmn" style="display: none;">
</ul>
<script type="text/javascript" reload="1">checkmgcmn('post_19135338')</script>
</div>
</div>
</td>
</tr>
<tr class="ad">
<td class="pls">
</td>
<td class="plc">
</td>
</tr>
</table>
</div>

<div id="post_19135551" ><table id="pid19135551" class="plhin" summary="pid19135551" cellspacing="0" cellpadding="0">
<tr>
 <td class="pls" rowspan="2">
<div id="favatar19135551" class="pls cl favatar">
<div class="pi">
<div class="authi"><a href="space-uid-483833.html" target="_blank" class="xw1">ycrepycrep</a>
</div>
</div>
<div class="p_pop blk bui card_gender_0" id="userinfo19135551" style="display: none; margin-top: -11px;">
<div class="m z">
<div id="userinfo19135551_ma"></div>
</div>
<div class="i y">
<div>
<strong><a href="space-uid-483833.html" target="_blank" class="xi2">ycrepycrep</a></strong>
<em>当前在线</em>
</div><dl class="cl">
<dt>积分</dt><dd><a href="home.php?mod=space&uid=483833&do=profile" target="_blank" class="xi2">58</a></dd>
</dl><div class="imicn">
<a href="home.php?mod=space&amp;uid=483833&amp;do=profile" target="_blank" title="查看详细资料"><img src="static/image/common/userinfo.gif" alt="查看详细资料" /></a>
</div>
<div id="avatarfeed"><span id="threadsortswait"></span></div>
</div>
</div>
<div>
<div class="avatar" onmouseover="showauthor(this, 'userinfo19135551')"><a href="space-uid-483833.html" class="avtm" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/48/38/33_avatar_middle.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_middle.gif'" /></a></div>
</div>
<div class="tns xg2"><table cellspacing="0" cellpadding="0"><th><p><a href="home.php?mod=space&uid=483833&do=profile" class="xi2">12</a></p>金钱</th><th><p><a href="home.php?mod=space&uid=483833&do=profile" class="xi2">0</a></p>色币</th><td><p><a href="home.php?mod=space&uid=483833&do=profile" class="xi2">0</a></p>评分</td></table></div>

<p><em><a href="home.php?mod=spacecp&amp;ac=usergroup&amp;gid=13" target="_blank">Lv4 锋芒毕露</a></em></p>


<p><span class="pbg2"  id="upgradeprogress_19135551" onmouseover="showMenu({'ctrlid':this.id, 'pos':'12!', 'menuid':'g_up19135551_menu'});"><span class="pbr2" style="width:16%;"></span></span></p>
<div id="g_up19135551_menu" class="tip tip_4" style="display: none;"><div class="tip_horn"></div><div class="tip_c">Lv4 锋芒毕露, 积分 58, 距离下一级还需 42 积分</div></div>

<dl class="pil cl">
  <dt>积分</dt><dd><a href="home.php?mod=space&uid=483833&do=profile" target="_blank" class="xi2">58</a></dd>
</dl>


<ul class="xl xl2 o cl">
<li class="pm2"><a href="home.php?mod=spacecp&amp;ac=pm&amp;op=showmsg&amp;handlekey=showmsg_483833&amp;touid=483833&amp;pmid=0&amp;daterange=2&amp;pid=19135551&amp;tid=1901572" onclick="showWindow('sendpm', this.href);" title="发消息" class="xi2">发消息</a></li>
</ul>
</div>
</td>
<td class="plc">
<div class="pi">
<strong>
<a href="forum.php?mod=redirect&goto=findpost&ptid=1901572&pid=19135551"   id="postnum19135551" onclick="setCopy(this.href, '帖子地址复制成功');return false;">
<em>5</em><sup>#</sup></a>
</strong>
<div class="pti">
<div class="pdbt">
</div>
<div class="authi">
<img class="authicn vm" id="authicon19135551" src="static/image/common/online_member.gif" />
<em id="authorposton19135551">发表于 <span title="2024-03-02 09:26:22">昨天&nbsp;09:26</span></em>
<span class="xg1">来自手机</span>
<span class="pipe">|</span>
<a href="forum.php?mod=viewthread&amp;tid=1901572&amp;page=1&amp;authorid=483833" rel="nofollow">只看该作者</a>
</div>
</div>
</div><div class="pct"><div class="pcb">
<div class="t_fsz">
<table cellspacing="0" cellpadding="0"><tr><td class="t_f" id="postmessage_19135551">
挺不错的，感谢</td></tr></table>


</div>
<div id="comment_19135551" class="cm">
</div>

<div id="post_rate_div_19135551"></div>
</div>
</div>

</td></tr>
<tr><td class="plc plm">

<div id="p_btn" class="mtw mbm hm cl">
</div>

</td>
</tr>
<tr id="_postposition19135551"></tr>
<tr>
<td class="pls"></td>
<td class="plc" style="overflow:visible;">
<div class="po hin">
<div class="pob cl">
<em>
<a class="fastre" href="forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1901572&amp;repquote=19135551&amp;extra=page%3D1&amp;page=1" onclick="showWindow('reply', this.href)">回复</a>
</em>

<p>
<a href="javascript:;" id="mgc_post_19135551" onmouseover="showMenu(this.id)" class="showmenu">使用道具</a>
<a href="javascript:;" onclick="showWindow('miscreport19135551', 'misc.php?mod=report&rtype=post&rid=19135551&tid=1901572&fid=36', 'get', -1);return false;">举报</a>
</p>

<ul id="mgc_post_19135551_menu" class="p_pop mgcmn" style="display: none;">
</ul>
<script type="text/javascript" reload="1">checkmgcmn('post_19135551')</script>
</div>
</div>
</td>
</tr>
<tr class="ad">
<td class="pls">
</td>
<td class="plc">
</td>
</tr>
</table>
</div>

<div id="post_19135745" ><table id="pid19135745" class="plhin" summary="pid19135745" cellspacing="0" cellpadding="0">
<tr>
 <td class="pls" rowspan="2">
<div id="favatar19135745" class="pls cl favatar">
<div class="pi">
<div class="authi"><a href="space-uid-470153.html" target="_blank" class="xw1">zh1993</a>
</div>
</div>
<div class="p_pop blk bui card_gender_0" id="userinfo19135745" style="display: none; margin-top: -11px;">
<div class="m z">
<div id="userinfo19135745_ma"></div>
</div>
<div class="i y">
<div>
<strong><a href="space-uid-470153.html" target="_blank" class="xi2">zh1993</a></strong>
<em>当前离线</em>
</div><dl class="cl">
<dt>积分</dt><dd><a href="home.php?mod=space&uid=470153&do=profile" target="_blank" class="xi2">113</a></dd>
</dl><div class="imicn">
<a href="home.php?mod=space&amp;uid=470153&amp;do=profile" target="_blank" title="查看详细资料"><img src="static/image/common/userinfo.gif" alt="查看详细资料" /></a>
</div>
<div id="avatarfeed"><span id="threadsortswait"></span></div>
</div>
</div>
<div>
<div class="avatar" onmouseover="showauthor(this, 'userinfo19135745')"><a href="space-uid-470153.html" class="avtm" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/47/01/53_avatar_middle.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_middle.gif'" /></a></div>
</div>
<div class="tns xg2"><table cellspacing="0" cellpadding="0"><th><p><a href="home.php?mod=space&uid=470153&do=profile" class="xi2">242</a></p>金钱</th><th><p><a href="home.php?mod=space&uid=470153&do=profile" class="xi2">0</a></p>色币</th><td><p><a href="home.php?mod=space&uid=470153&do=profile" class="xi2">10</a></p>评分</td></table></div>

<p><em><a href="home.php?mod=spacecp&amp;ac=usergroup&amp;gid=14" target="_blank">Lv5 小有名气</a></em></p>


<p><span class="pbg2"  id="upgradeprogress_19135745" onmouseover="showMenu({'ctrlid':this.id, 'pos':'12!', 'menuid':'g_up19135745_menu'});"><span class="pbr2" style="width:13%;"></span></span></p>
<div id="g_up19135745_menu" class="tip tip_4" style="display: none;"><div class="tip_horn"></div><div class="tip_c">Lv5 小有名气, 积分 113, 距离下一级还需 87 积分</div></div>

<dl class="pil cl">
  <dt>积分</dt><dd><a href="home.php?mod=space&uid=470153&do=profile" target="_blank" class="xi2">113</a></dd>
</dl>


<ul class="xl xl2 o cl">
<li class="pm2"><a href="home.php?mod=spacecp&amp;ac=pm&amp;op=showmsg&amp;handlekey=showmsg_470153&amp;touid=470153&amp;pmid=0&amp;daterange=2&amp;pid=19135745&amp;tid=1901572" onclick="showWindow('sendpm', this.href);" title="发消息" class="xi2">发消息</a></li>
</ul>
</div>
</td>
<td class="plc">
<div class="pi">
<strong>
<a href="forum.php?mod=redirect&goto=findpost&ptid=1901572&pid=19135745"   id="postnum19135745" onclick="setCopy(this.href, '帖子地址复制成功');return false;">
<em>6</em><sup>#</sup></a>
</strong>
<div class="pti">
<div class="pdbt">
</div>
<div class="authi">
<img class="authicn vm" id="authicon19135745" src="static/image/common/online_member.gif" />
<em id="authorposton19135745">发表于 <span title="2024-03-02 09:32:53">昨天&nbsp;09:32</span></em>
<span class="pipe">|</span>
<a href="forum.php?mod=viewthread&amp;tid=1901572&amp;page=1&amp;authorid=470153" rel="nofollow">只看该作者</a>
</div>
</div>
</div><div class="pct"><div class="pcb">
<div class="t_fsz">
<table cellspacing="0" cellpadding="0"><tr><td class="t_f" id="postmessage_19135745">
今日最佳</td></tr></table>


</div>
<div id="comment_19135745" class="cm">
</div>

<div id="post_rate_div_19135745"></div>
</div>
</div>

</td></tr>
<tr><td class="plc plm">

<div id="p_btn" class="mtw mbm hm cl">
</div>

</td>
</tr>
<tr id="_postposition19135745"></tr>
<tr>
<td class="pls"></td>
<td class="plc" style="overflow:visible;">
<div class="po hin">
<div class="pob cl">
<em>
<a class="fastre" href="forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1901572&amp;repquote=19135745&amp;extra=page%3D1&amp;page=1" onclick="showWindow('reply', this.href)">回复</a>
</em>

<p>
<a href="javascript:;" id="mgc_post_19135745" onmouseover="showMenu(this.id)" class="showmenu">使用道具</a>
<a href="javascript:;" onclick="showWindow('miscreport19135745', 'misc.php?mod=report&rtype=post&rid=19135745&tid=1901572&fid=36', 'get', -1);return false;">举报</a>
</p>

<ul id="mgc_post_19135745_menu" class="p_pop mgcmn" style="display: none;">
</ul>
<script type="text/javascript" reload="1">checkmgcmn('post_19135745')</script>
</div>
</div>
</td>
</tr>
<tr class="ad">
<td class="pls">
</td>
<td class="plc">
</td>
</tr>
</table>
</div>

<div id="post_19136497" ><table id="pid19136497" class="plhin" summary="pid19136497" cellspacing="0" cellpadding="0">
<tr>
 <td class="pls" rowspan="2">
<div id="favatar19136497" class="pls cl favatar">
<div class="pi">
<div class="authi"><a href="space-uid-433969.html" target="_blank" class="xw1">tony1985</a>
</div>
</div>
<div class="p_pop blk bui card_gender_0" id="userinfo19136497" style="display: none; margin-top: -11px;">
<div class="m z">
<div id="userinfo19136497_ma"></div>
</div>
<div class="i y">
<div>
<strong><a href="space-uid-433969.html" target="_blank" class="xi2">tony1985</a></strong>
<em>当前在线</em>
</div><dl class="cl">
<dt>积分</dt><dd><a href="home.php?mod=space&uid=433969&do=profile" target="_blank" class="xi2">203</a></dd>
</dl><div class="imicn">
<a href="home.php?mod=space&amp;uid=433969&amp;do=profile" target="_blank" title="查看详细资料"><img src="static/image/common/userinfo.gif" alt="查看详细资料" /></a>
</div>
<div id="avatarfeed"><span id="threadsortswait"></span></div>
</div>
</div>
<div>
<div class="avatar" onmouseover="showauthor(this, 'userinfo19136497')"><a href="space-uid-433969.html" class="avtm" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/43/39/69_avatar_middle.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_middle.gif'" /></a></div>
</div>
<div class="tns xg2"><table cellspacing="0" cellpadding="0"><th><p><a href="home.php?mod=space&uid=433969&do=profile" class="xi2">153</a></p>金钱</th><th><p><a href="home.php?mod=space&uid=433969&do=profile" class="xi2">0</a></p>色币</th><td><p><a href="home.php?mod=space&uid=433969&do=profile" class="xi2">17</a></p>评分</td></table></div>

<p><em><a href="home.php?mod=spacecp&amp;ac=usergroup&amp;gid=15" target="_blank">Lv6 后起之秀</a></em></p>


<p><span class="pbg2"  id="upgradeprogress_19136497" onmouseover="showMenu({'ctrlid':this.id, 'pos':'12!', 'menuid':'g_up19136497_menu'});"><span class="pbr2" style="width:2%;"></span></span></p>
<div id="g_up19136497_menu" class="tip tip_4" style="display: none;"><div class="tip_horn"></div><div class="tip_c">Lv6 后起之秀, 积分 203, 距离下一级还需 297 积分</div></div>

<dl class="pil cl">
  <dt>积分</dt><dd><a href="home.php?mod=space&uid=433969&do=profile" target="_blank" class="xi2">203</a></dd>
</dl>


<ul class="xl xl2 o cl">
<li class="pm2"><a href="home.php?mod=spacecp&amp;ac=pm&amp;op=showmsg&amp;handlekey=showmsg_433969&amp;touid=433969&amp;pmid=0&amp;daterange=2&amp;pid=19136497&amp;tid=1901572" onclick="showWindow('sendpm', this.href);" title="发消息" class="xi2">发消息</a></li>
</ul>
</div>
</td>
<td class="plc">
<div class="pi">
<strong>
<a href="forum.php?mod=redirect&goto=findpost&ptid=1901572&pid=19136497"   id="postnum19136497" onclick="setCopy(this.href, '帖子地址复制成功');return false;">
<em>7</em><sup>#</sup></a>
</strong>
<div class="pti">
<div class="pdbt">
</div>
<div class="authi">
<img class="authicn vm" id="authicon19136497" src="static/image/common/online_member.gif" />
<em id="authorposton19136497">发表于 <span title="2024-03-02 09:57:15">昨天&nbsp;09:57</span></em>
<span class="pipe">|</span>
<a href="forum.php?mod=viewthread&amp;tid=1901572&amp;page=1&amp;authorid=433969" rel="nofollow">只看该作者</a>
</div>
</div>
</div><div class="pct"><div class="pcb">
<div class="t_fsz">
<table cellspacing="0" cellpadding="0"><tr><td class="t_f" id="postmessage_19136497">
<br />
感谢分享</td></tr></table>


</div>
<div id="comment_19136497" class="cm">
</div>

<div id="post_rate_div_19136497"></div>
</div>
</div>

</td></tr>
<tr><td class="plc plm">

<div id="p_btn" class="mtw mbm hm cl">
</div>

</td>
</tr>
<tr id="_postposition19136497"></tr>
<tr>
<td class="pls"></td>
<td class="plc" style="overflow:visible;">
<div class="po hin">
<div class="pob cl">
<em>
<a class="fastre" href="forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1901572&amp;repquote=19136497&amp;extra=page%3D1&amp;page=1" onclick="showWindow('reply', this.href)">回复</a>
</em>

<p>
<a href="javascript:;" id="mgc_post_19136497" onmouseover="showMenu(this.id)" class="showmenu">使用道具</a>
<a href="javascript:;" onclick="showWindow('miscreport19136497', 'misc.php?mod=report&rtype=post&rid=19136497&tid=1901572&fid=36', 'get', -1);return false;">举报</a>
</p>

<ul id="mgc_post_19136497_menu" class="p_pop mgcmn" style="display: none;">
</ul>
<script type="text/javascript" reload="1">checkmgcmn('post_19136497')</script>
</div>
</div>
</td>
</tr>
<tr class="ad">
<td class="pls">
</td>
<td class="plc">
</td>
</tr>
</table>
</div>

<div id="post_19136987" ><table id="pid19136987" class="plhin" summary="pid19136987" cellspacing="0" cellpadding="0">
<tr>
 <td class="pls" rowspan="2">
<div id="favatar19136987" class="pls cl favatar">
<div class="pi">
<div class="authi"><a href="space-uid-468504.html" target="_blank" class="xw1">欧金兴</a>
</div>
</div>
<div class="p_pop blk bui card_gender_0" id="userinfo19136987" style="display: none; margin-top: -11px;">
<div class="m z">
<div id="userinfo19136987_ma"></div>
</div>
<div class="i y">
<div>
<strong><a href="space-uid-468504.html" target="_blank" class="xi2">欧金兴</a></strong>
<em>当前在线</em>
</div><dl class="cl">
<dt>积分</dt><dd><a href="home.php?mod=space&uid=468504&do=profile" target="_blank" class="xi2">1098</a></dd>
</dl><div class="imicn">
<a href="home.php?mod=space&amp;uid=468504&amp;do=profile" target="_blank" title="查看详细资料"><img src="static/image/common/userinfo.gif" alt="查看详细资料" /></a>
</div>
<div id="avatarfeed"><span id="threadsortswait"></span></div>
</div>
</div>
<div>
<div class="avatar" onmouseover="showauthor(this, 'userinfo19136987')"><a href="space-uid-468504.html" class="avtm" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/46/85/04_avatar_middle.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_middle.gif'" /></a></div>
</div>
<div class="tns xg2"><table cellspacing="0" cellpadding="0"><th><p><a href="home.php?mod=space&uid=468504&do=profile" class="xi2">267</a></p>金钱</th><th><p><a href="home.php?mod=space&uid=468504&do=profile" class="xi2">0</a></p>色币</th><td><p><a href="home.php?mod=space&uid=468504&do=profile" class="xi2">0</a></p>评分</td></table></div>

<p><em><a href="home.php?mod=spacecp&amp;ac=usergroup&amp;gid=28" target="_blank">Lv8 武林高手</a></em></p>


<p><span class="pbg2"  id="upgradeprogress_19136987" onmouseover="showMenu({'ctrlid':this.id, 'pos':'12!', 'menuid':'g_up19136987_menu'});"><span class="pbr2" style="width:9%;"></span></span></p>
<div id="g_up19136987_menu" class="tip tip_4" style="display: none;"><div class="tip_horn"></div><div class="tip_c">Lv8 武林高手, 积分 1098, 距离下一级还需 902 积分</div></div>

<dl class="pil cl">
  <dt>积分</dt><dd><a href="home.php?mod=space&uid=468504&do=profile" target="_blank" class="xi2">1098</a></dd>
</dl>


<ul class="xl xl2 o cl">
<li class="pm2"><a href="home.php?mod=spacecp&amp;ac=pm&amp;op=showmsg&amp;handlekey=showmsg_468504&amp;touid=468504&amp;pmid=0&amp;daterange=2&amp;pid=19136987&amp;tid=1901572" onclick="showWindow('sendpm', this.href);" title="发消息" class="xi2">发消息</a></li>
</ul>
</div>
</td>
<td class="plc">
<div class="pi">
<strong>
<a href="forum.php?mod=redirect&goto=findpost&ptid=1901572&pid=19136987"   id="postnum19136987" onclick="setCopy(this.href, '帖子地址复制成功');return false;">
<em>8</em><sup>#</sup></a>
</strong>
<div class="pti">
<div class="pdbt">
</div>
<div class="authi">
<img class="authicn vm" id="authicon19136987" src="static/image/common/online_member.gif" />
<em id="authorposton19136987">发表于 <span title="2024-03-02 10:13:59">昨天&nbsp;10:13</span></em>
<span class="xg1">来自手机</span>
<span class="pipe">|</span>
<a href="forum.php?mod=viewthread&amp;tid=1901572&amp;page=1&amp;authorid=468504" rel="nofollow">只看该作者</a>
</div>
</div>
</div><div class="pct"><div class="pcb">
<div class="t_fsz">
<table cellspacing="0" cellpadding="0"><tr><td class="t_f" id="postmessage_19136987">
近期来说很赞了，身材丰满</td></tr></table>


</div>
<div id="comment_19136987" class="cm">
</div>

<div id="post_rate_div_19136987"></div>
</div>
</div>

</td></tr>
<tr><td class="plc plm">

<div id="p_btn" class="mtw mbm hm cl">
</div>

</td>
</tr>
<tr id="_postposition19136987"></tr>
<tr>
<td class="pls"></td>
<td class="plc" style="overflow:visible;">
<div class="po hin">
<div class="pob cl">
<em>
<a class="fastre" href="forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1901572&amp;repquote=19136987&amp;extra=page%3D1&amp;page=1" onclick="showWindow('reply', this.href)">回复</a>
</em>

<p>
<a href="javascript:;" id="mgc_post_19136987" onmouseover="showMenu(this.id)" class="showmenu">使用道具</a>
<a href="javascript:;" onclick="showWindow('miscreport19136987', 'misc.php?mod=report&rtype=post&rid=19136987&tid=1901572&fid=36', 'get', -1);return false;">举报</a>
</p>

<ul id="mgc_post_19136987_menu" class="p_pop mgcmn" style="display: none;">
</ul>
<script type="text/javascript" reload="1">checkmgcmn('post_19136987')</script>
</div>
</div>
</td>
</tr>
<tr class="ad">
<td class="pls">
</td>
<td class="plc">
</td>
</tr>
</table>
</div>

<div id="post_19137064" ><table id="pid19137064" class="plhin" summary="pid19137064" cellspacing="0" cellpadding="0">
<tr>
 <td class="pls" rowspan="2">
<div id="favatar19137064" class="pls cl favatar">
<div class="pi">
<div class="authi"><a href="space-uid-471291.html" target="_blank" class="xw1">lemon000</a>
</div>
</div>
<div class="p_pop blk bui card_gender_0" id="userinfo19137064" style="display: none; margin-top: -11px;">
<div class="m z">
<div id="userinfo19137064_ma"></div>
</div>
<div class="i y">
<div>
<strong><a href="space-uid-471291.html" target="_blank" class="xi2">lemon000</a></strong>
<em>当前离线</em>
</div><dl class="cl">
<dt>积分</dt><dd><a href="home.php?mod=space&uid=471291&do=profile" target="_blank" class="xi2">381</a></dd>
</dl><div class="imicn">
<a href="home.php?mod=space&amp;uid=471291&amp;do=profile" target="_blank" title="查看详细资料"><img src="static/image/common/userinfo.gif" alt="查看详细资料" /></a>
</div>
<div id="avatarfeed"><span id="threadsortswait"></span></div>
</div>
</div>
<div>
<div class="avatar" onmouseover="showauthor(this, 'userinfo19137064')"><a href="space-uid-471291.html" class="avtm" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/47/12/91_avatar_middle.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_middle.gif'" /></a></div>
</div>
<div class="tns xg2"><table cellspacing="0" cellpadding="0"><th><p><a href="home.php?mod=space&uid=471291&do=profile" class="xi2">215</a></p>金钱</th><th><p><a href="home.php?mod=space&uid=471291&do=profile" class="xi2">0</a></p>色币</th><td><p><a href="home.php?mod=space&uid=471291&do=profile" class="xi2">4</a></p>评分</td></table></div>

<p><em><a href="home.php?mod=spacecp&amp;ac=usergroup&amp;gid=15" target="_blank">Lv6 后起之秀</a></em></p>


<p><span class="pbg2"  id="upgradeprogress_19137064" onmouseover="showMenu({'ctrlid':this.id, 'pos':'12!', 'menuid':'g_up19137064_menu'});"><span class="pbr2" style="width:60%;"></span></span></p>
<div id="g_up19137064_menu" class="tip tip_4" style="display: none;"><div class="tip_horn"></div><div class="tip_c">Lv6 后起之秀, 积分 381, 距离下一级还需 119 积分</div></div>

<dl class="pil cl">
  <dt>积分</dt><dd><a href="home.php?mod=space&uid=471291&do=profile" target="_blank" class="xi2">381</a></dd>
</dl>


<ul class="xl xl2 o cl">
<li class="pm2"><a href="home.php?mod=spacecp&amp;ac=pm&amp;op=showmsg&amp;handlekey=showmsg_471291&amp;touid=471291&amp;pmid=0&amp;daterange=2&amp;pid=19137064&amp;tid=1901572" onclick="showWindow('sendpm', this.href);" title="发消息" class="xi2">发消息</a></li>
</ul>
</div>
</td>
<td class="plc">
<div class="pi">
<strong>
<a href="forum.php?mod=redirect&goto=findpost&ptid=1901572&pid=19137064"   id="postnum19137064" onclick="setCopy(this.href, '帖子地址复制成功');return false;">
<em>9</em><sup>#</sup></a>
</strong>
<div class="pti">
<div class="pdbt">
</div>
<div class="authi">
<img class="authicn vm" id="authicon19137064" src="static/image/common/online_member.gif" />
<em id="authorposton19137064">发表于 <span title="2024-03-02 10:16:52">昨天&nbsp;10:16</span></em>
<span class="pipe">|</span>
<a href="forum.php?mod=viewthread&amp;tid=1901572&amp;page=1&amp;authorid=471291" rel="nofollow">只看该作者</a>
</div>
</div>
</div><div class="pct"><div class="pcb">
<div class="t_fsz">
<table cellspacing="0" cellpadding="0"><tr><td class="t_f" id="postmessage_19137064">
<br />
感谢分享</td></tr></table>


</div>
<div id="comment_19137064" class="cm">
</div>

<div id="post_rate_div_19137064"></div>
</div>
</div>

</td></tr>
<tr><td class="plc plm">

<div id="p_btn" class="mtw mbm hm cl">
</div>

</td>
</tr>
<tr id="_postposition19137064"></tr>
<tr>
<td class="pls"></td>
<td class="plc" style="overflow:visible;">
<div class="po hin">
<div class="pob cl">
<em>
<a class="fastre" href="forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1901572&amp;repquote=19137064&amp;extra=page%3D1&amp;page=1" onclick="showWindow('reply', this.href)">回复</a>
</em>

<p>
<a href="javascript:;" id="mgc_post_19137064" onmouseover="showMenu(this.id)" class="showmenu">使用道具</a>
<a href="javascript:;" onclick="showWindow('miscreport19137064', 'misc.php?mod=report&rtype=post&rid=19137064&tid=1901572&fid=36', 'get', -1);return false;">举报</a>
</p>

<ul id="mgc_post_19137064_menu" class="p_pop mgcmn" style="display: none;">
</ul>
<script type="text/javascript" reload="1">checkmgcmn('post_19137064')</script>
</div>
</div>
</td>
</tr>
<tr class="ad">
<td class="pls">
</td>
<td class="plc">
</td>
</tr>
</table>
</div>

<div id="post_19138329" ><table id="pid19138329" class="plhin" summary="pid19138329" cellspacing="0" cellpadding="0">
<tr>
 <td class="pls" rowspan="2">
<div id="favatar19138329" class="pls cl favatar">
<div class="pi">
<div class="authi"><a href="space-uid-456768.html" target="_blank" class="xw1">bayxsy0124</a>
</div>
</div>
<div class="p_pop blk bui card_gender_0" id="userinfo19138329" style="display: none; margin-top: -11px;">
<div class="m z">
<div id="userinfo19138329_ma"></div>
</div>
<div class="i y">
<div>
<strong><a href="space-uid-456768.html" target="_blank" class="xi2">bayxsy0124</a></strong>
<em>当前离线</em>
</div><dl class="cl">
<dt>积分</dt><dd><a href="home.php?mod=space&uid=456768&do=profile" target="_blank" class="xi2">829</a></dd>
</dl><div class="imicn">
<a href="home.php?mod=space&amp;uid=456768&amp;do=profile" target="_blank" title="查看详细资料"><img src="static/image/common/userinfo.gif" alt="查看详细资料" /></a>
</div>
<div id="avatarfeed"><span id="threadsortswait"></span></div>
</div>
</div>
<div>
<div class="avatar" onmouseover="showauthor(this, 'userinfo19138329')"><a href="space-uid-456768.html" class="avtm" target="_blank"><img src="https://qz1gy.app/uc_server/data/avatar/000/45/67/68_avatar_middle.jpg" onerror="this.onerror=null;this.src='https://qz1gy.app/uc_server/images/noavatar_middle.gif'" /></a></div>
</div>
<div class="tns xg2"><table cellspacing="0" cellpadding="0"><th><p><a href="home.php?mod=space&uid=456768&do=profile" class="xi2">428</a></p>金钱</th><th><p><a href="home.php?mod=space&uid=456768&do=profile" class="xi2">0</a></p>色币</th><td><p><a href="home.php?mod=space&uid=456768&do=profile" class="xi2">0</a></p>评分</td></table></div>

<p><em><a href="home.php?mod=spacecp&amp;ac=usergroup&amp;gid=27" target="_blank">Lv7 四方游侠</a></em></p>


<p><span class="pbg2"  id="upgradeprogress_19138329" onmouseover="showMenu({'ctrlid':this.id, 'pos':'12!', 'menuid':'g_up19138329_menu'});"><span class="pbr2" style="width:65%;"></span></span></p>
<div id="g_up19138329_menu" class="tip tip_4" style="display: none;"><div class="tip_horn"></div><div class="tip_c">Lv7 四方游侠, 积分 829, 距离下一级还需 171 积分</div></div>

<dl class="pil cl">
  <dt>积分</dt><dd><a href="home.php?mod=space&uid=456768&do=profile" target="_blank" class="xi2">829</a></dd>
</dl>


<ul class="xl xl2 o cl">
<li class="pm2"><a href="home.php?mod=spacecp&amp;ac=pm&amp;op=showmsg&amp;handlekey=showmsg_456768&amp;touid=456768&amp;pmid=0&amp;daterange=2&amp;pid=19138329&amp;tid=1901572" onclick="showWindow('sendpm', this.href);" title="发消息" class="xi2">发消息</a></li>
</ul>
</div>
</td>
<td class="plc">
<div class="pi">
<strong>
<a href="forum.php?mod=redirect&goto=findpost&ptid=1901572&pid=19138329"   id="postnum19138329" onclick="setCopy(this.href, '帖子地址复制成功');return false;">
<em>10</em><sup>#</sup></a>
</strong>
<div class="pti">
<div class="pdbt">
</div>
<div class="authi">
<img class="authicn vm" id="authicon19138329" src="static/image/common/online_member.gif" />
<em id="authorposton19138329">发表于 <span title="2024-03-02 10:57:07">昨天&nbsp;10:57</span></em>
<span class="pipe">|</span>
<a href="forum.php?mod=viewthread&amp;tid=1901572&amp;page=1&amp;authorid=456768" rel="nofollow">只看该作者</a>
</div>
</div>
</div><div class="pct"><div class="pcb">
<div class="t_fsz">
<table cellspacing="0" cellpadding="0"><tr><td class="t_f" id="postmessage_19138329">
不错，感谢分享</td></tr></table>


</div>
<div id="comment_19138329" class="cm">
</div>

<div id="post_rate_div_19138329"></div>
</div>
</div>

</td></tr>
<tr><td class="plc plm">

<div id="p_btn" class="mtw mbm hm cl">
</div>

</td>
</tr>
<tr id="_postposition19138329"></tr>
<tr>
<td class="pls"></td>
<td class="plc" style="overflow:visible;">
<div class="po hin">
<div class="pob cl">
<em>
<a class="fastre" href="forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1901572&amp;repquote=19138329&amp;extra=page%3D1&amp;page=1" onclick="showWindow('reply', this.href)">回复</a>
</em>

<p>
<a href="javascript:;" id="mgc_post_19138329" onmouseover="showMenu(this.id)" class="showmenu">使用道具</a>
<a href="javascript:;" onclick="showWindow('miscreport19138329', 'misc.php?mod=report&rtype=post&rid=19138329&tid=1901572&fid=36', 'get', -1);return false;">举报</a>
</p>

<ul id="mgc_post_19138329_menu" class="p_pop mgcmn" style="display: none;">
</ul>
<script type="text/javascript" reload="1">checkmgcmn('post_19138329')</script>
</div>
</div>
</td>
</tr>
<tr class="ad">
<td class="pls">
</td>
<td class="plc">
</td>
</tr>
</table>
</div>

<div id="postlistreply" class="pl"><div id="post_new" class="viewthread_table" style="display: none"></div></div>
</div>


<form method="post" autocomplete="off" name="modactions" id="modactions">
<input type="hidden" name="formhash" value="f472fc1f" />
<input type="hidden" name="optgroup" />
<input type="hidden" name="operation" />
<input type="hidden" name="listextra" value="page%3D1" />
<input type="hidden" name="page" value="1" />
</form>


<div class="pgbtn"><a href="thread-1901572-2-1.html" hidefocus="true" class="bm_h">下一页 &raquo;</a></div>

<div class="pgs mtm mbm cl">
<div class="pg"><strong>1</strong><a href="thread-1901572-2-1.html">2</a><a href="thread-1901572-3-1.html">3</a><a href="thread-1901572-4-1.html">4</a><a href="thread-1901572-5-1.html">5</a><a href="thread-1901572-6-1.html">6</a><a href="thread-1901572-7-1.html">7</a><label><input type="text" name="custompage" class="px" size="2" title="输入页码，按回车快速跳转" value="1" onkeydown="if(event.keyCode==13) {window.location='forum.php?mod=viewthread&tid=1901572&amp;extra=page%3D1&amp;page='+this.value;; doane(event);}" /><span title="共 7 页"> / 7 页</span></label><a href="thread-1901572-2-1.html" class="nxt">下一页</a></div><span class="pgb y" id="visitedforumstmp" onmouseover="$('visitedforums').id = 'visitedforumstmp';this.id = 'visitedforums';showMenu({'ctrlid':this.id,'pos':'21'})"><a href="forum-36-1.html">返回列表</a></span>
<a id="newspecialtmp" onmouseover="$('newspecial').id = 'newspecialtmp';this.id = 'newspecial';showMenu({'ctrlid':this.id})" onclick="showWindow('newthread', 'forum.php?mod=post&action=newthread&fid=36')" href="javascript:;" title="发新帖"><img src="static/image/common/pn_post.png" alt="发新帖" /></a>
</div>

<!--[diy=diyfastposttop]--><div id="diyfastposttop" class="area"></div><!--[/diy]-->
<script type="text/javascript">
var postminchars = parseInt('10');
var postmaxchars = parseInt('500000000');
var disablepostctrl = parseInt('0');
</script>

<div id="f_pst" class="pl bm bmw">
<form method="post" autocomplete="off" id="fastpostform" action="forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1901572&amp;extra=page%3D1&amp;replysubmit=yes&amp;infloat=yes&amp;handlekey=fastpost" onSubmit="return fastpostvalidate(this)">
<table cellspacing="0" cellpadding="0">
<tr>
<td class="pls">
</td>
<td class="plc">

<span id="fastpostreturn"></span>


<div class="cl">
<div id="fastsmiliesdiv" class="y"><div id="fastsmiliesdiv_data"><div id="fastsmilies"></div></div></div><div class="hasfsl" id="fastposteditor">
<div class="tedt mtn">
<div class="bar">
<span class="y">
<a href="forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1901572" onclick="return switchAdvanceMode(this.href)">高级模式</a>
</span><script src="static/js/seditor.js?OfY" type="text/javascript"></script>
<div class="fpd">
<a href="javascript:;" title="文字加粗" class="fbld">B</a>
<a href="javascript:;" title="设置文字颜色" class="fclr" id="fastpostforecolor">Color</a>
<a id="fastpostimg" href="javascript:;" title="图片" class="fmg">Image</a>
<a id="fastposturl" href="javascript:;" title="添加链接" class="flnk">Link</a>
<a id="fastpostquote" href="javascript:;" title="引用" class="fqt">Quote</a>
<a id="fastpostcode" href="javascript:;" title="代码" class="fcd">Code</a>
<a href="javascript:;" class="fsml" id="fastpostsml">Smilies</a>
</div></div>
<div class="area">
<div class="pt hm">
您需要登录后才可以回帖 <a href="member.php?mod=logging&amp;action=login" onclick="showWindow('login', this.href)" class="xi2">登录</a> | <a href="member.php?mod=register" class="xi2">立即注册</a>
</div>
</div>
</div>
</div>
</div>
<div id="seccheck_fastpost">
</div>


<input type="hidden" name="formhash" value="f472fc1f" />
<input type="hidden" name="usesig" value="" />
<input type="hidden" name="subject" value="  " />
<p class="ptm pnpost">
<a href="home.php?mod=spacecp&amp;ac=credit&amp;op=rule&amp;fid=36" class="y" target="_blank">本版积分规则</a>
<button type="button" onclick="showWindow('login', 'member.php?mod=logging&action=login&guestmessage=yes')"  onmouseover="checkpostrule('seccheck_fastpost', 'ac=reply');this.onmouseover=null" name="replysubmit" id="fastpostsubmit" class="pn pnc vm" value="replysubmit" tabindex="5"><strong>发表回复</strong></button>
<label for="fastpostrefresh"><input id="fastpostrefresh" type="checkbox" class="pc" />回帖后跳转到最后一页</label>
<script type="text/javascript">if(getcookie('fastpostrefresh') == 1) {$('fastpostrefresh').checked=true;}</script>
</p>
</td>
</tr>
</table>
</form>
</div>

<div id="visitedforums_menu" class="p_pop blk cl" style="display: none;">
<table cellspacing="0" cellpadding="0">
<tr>
<td id="v_forums">
<h3 class="mbn pbn bbda xg1">浏览过的版块</h3>
<ul class="xl xl1">
<li><a href="forum-104-1.html">素人有码系列</a></li><li><a href="forum-37-1.html">亚洲有码原创</a></li><li><a href="forum-143-1.html">求片问答悬赏区</a></li><li><a href="forum-95-1.html">综合讨论区</a></li><li><a href="forum-107-1.html">三级写真</a></li><li><a href="forum-139-1.html">TXT小说下载</a></li><li><a href="forum-142-1.html">转帖交流区</a></li><li><a href="forum-159-1.html">新作区</a></li><li><a href="forum-151-1.html">4K原版</a></li></ul>
</td>
</tr>
</table>
</div>
<script type="text/javascript">
new lazyload();
</script>
<script type="text/javascript">document.onkeyup = function(e){keyPageScroll(e, 0, 1, 'forum.php?mod=viewthread&tid=1901572', 1);}</script>
</div>

<div class="wp mtn">
<!--[diy=diy3]--><div id="diy3" class="area"></div><!--[/diy]-->
</div>

<script type="text/javascript">
function succeedhandle_followmod(url, msg, values) {
var fObj = $('followmod_'+values['fuid']);
if(values['type'] == 'add') {
fObj.innerHTML = '不收听';
fObj.href = 'home.php?mod=spacecp&ac=follow&op=del&fuid='+values['fuid'];
} else if(values['type'] == 'del') {
fObj.innerHTML = '收听TA';
fObj.href = 'home.php?mod=spacecp&ac=follow&op=add&hash=a402aa39&fuid='+values['fuid'];
}
}
fixed_avatar([19133577,19133954,19134901,19135338,19135551,19135745,19136497,19136987,19137064,19138329], 1);
</script><div class="show-text cl"><a href="https://ls399.cc/?shareName=&proxyAccount=99482798&vertical=1#/" target="_blank" class="js-randomBg">拉斯维加斯</a>
<a href="https://hg9300f.cc:8989/?c=UNMHW" target="_blank" class="js-randomBg">澳门高爆电子</a>
<a href="https://ciudadpromo.com" target="_blank" class="js-randomBg">亚博美女赌场</a>
<a href="https://by301270224.cc/?channelCode=by_147" target="_blank" class="js-randomBg">鲍鱼全球黄播</a>
<a href="https://4qz.me" target="_blank" class="js-randomBg">免费成人乱伦</a>
<a href="https://572c.tihlrhpe.vip/aff-ucSu" target="_blank" class="js-randomBg">草榴成人免费</a>
<a href="https://cfslpis.wkvudxj.xyz/sht666-41" target="_blank" class="js-randomBg">免费萝莉禁区</a>
<a href="https://asfd6ef.com/yj/6475/zw98sh1awjq" target="_blank" class="js-randomBg">免费顶级暗网</a>
<a href="https://nwhtizm.wcerxcf.xyz/sht888-30" target="_blank" class="js-randomBg">91免费射区</a>
<a href="https://web.jbwa.ltd" target="_blank" class="js-randomBg">91成人抖音</a>
</div>
<div class="h10"></div>
	</div>
<div id="ft" class="wp cl">
<div id="flk" class="y">
<p>
广告联系：sehuatang@gmail.com<span class="pipe">|</span>
<a href="forum.php?mod=misc&action=showdarkroom" >小黑屋</a><span class="pipe">|</span><strong><a href="https://www.sehuatang.net" target="_blank">98堂[原色花堂]</a></strong>
</p>
<p class="xs0">
GMT+8, 2024-03-03 20:29<span id="debuginfo">
, Processed in 0.054815 second(s), 6 queries
, Gzip On, Redis On.
</span>
</p>
</div>
<div id="frt">
<p>Powered by <strong><a href="http://www.discuz.net" target="_blank">Discuz!</a></strong> <em>X3.4</em></p>
<p class="xs0">Copyright &copy; 2001-2021, Tencent Cloud.</p>
</div></div>

<div id="scrolltop">
<span><a href="forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1901572&amp;extra=page%3D1&amp;page=1" onclick="showWindow('reply', this.href)" class="replyfast" title="快速回复"><b>快速回复</b></a></span>
<span hidefocus="true"><a title="返回顶部" onclick="window.scrollTo('0','0')" class="scrolltopa" ><b>返回顶部</b></a></span>
<span>
<a href="forum-36-1.html" hidefocus="true" class="returnlist" title="返回列表"><b>返回列表</b></a>
</span>
</div>
<script type="text/javascript">_attachEvent(window, 'scroll', function () { showTopLink(); });checkBlind();randomBg('js-randomBg');</script>
<script type="text/javascript">$("debuginfo") ? $("debuginfo").innerHTML = ", Updated at 2024-03-03 20:29:16, Processed in 0.004177 second(s)." : "";</script></body></html>`
	xpath := "/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/div[@id]/table[@id]/tbody/tr[1]"
	//config := "{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\",\"obj\":{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\"}}"
	srt := ""

	resultHandle := "{\n    \"user\": {\n        \"href_attr\":\"/tr/td[@class=\\\"pls\\\"]/div[@class=\\\"pls cl favatar\\\"]/div[@class=\\\"pi\\\"]/div[@class=\\\"authi\\\"]/a\",\n        \"name_content\":\"/tr/td[@class=\\\"pls\\\"]/div[@class=\\\"pls cl favatar\\\"]/div[@class=\\\"pi\\\"]/div[@class=\\\"authi\\\"]/a\"\n    },\n    \"index_content\": \"/tr/td[@class=\\\"plc\\\"]/div[@class=\\\"pi\\\"]/strong/a\",\n    \"time\": {\n       \"title_attr\" :\"/tr/td[@class=\\\"plc\\\"]/div[@class=\\\"pi\\\"]/div[@class=\\\"pti\\\"]//em/span\"\n    },\n    \"evaluate_content\": \"/tr/td[@class=\\\"plc\\\"]/div[@class=\\\"pct\\\"]/div[@class=\\\"pcb\\\"]/div[@class=\\\"t_fsz\\\"]//td\"\n}"
	fmt.Println(Exe(srt, xpath, html, resultHandle))
}

func TestConfig(t *testing.T) {
	html := "<!DOCTYPE html> <html> <head> <meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" /> <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title> </body> </html>"
	//xpath := "//title"
	config := "{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\",\"obj\":{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\"}}"
	srt := ""
	fmt.Println(Exe(config, srt, html, ""))
}

func TestHtml(t *testing.T) {
	html := "<!DOCTYPE html>\r\n<html>\r\n<head>\r\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" />\r\n<title>SSNI-116  [无码破解] 絶頂してピクピクしているおま●こを容赦なく突きまくる怒涛の - 亚洲无码原创 -  98堂[原色花堂] -  Powered by Discuz!</title>\r\n<link href=\"/thread-1291126-1-1.html\" rel=\"canonical\" />\r\n<meta name=\"keywords\" content=\"SSNI-116  [无码破解] 絶頂してピクピクしているおま●こを容赦なく突きまくる怒涛の\" />\r\n<meta name=\"description\" content=\"【影片名称】：絶頂してピクピクしているおま●こを容赦なく突きまくる怒涛のおかわり激ピストン性交 【出演女优】：天使もえ【影片容量】：5.07G【是否有码】：无码【种子期限】：5种或健康度1000【下载工具】 ... SSNI-116  [无码破解] 絶頂してピクピクしているおま●こを容赦なく突きまくる怒涛の ,98堂[原色花堂]\" />\r\n<meta name=\"generator\" content=\"Discuz! X3.4\" />\r\n<meta name=\"author\" content=\"Discuz! Team and Comsenz UI Team\" />\r\n<meta name=\"copyright\" content=\"2001-2021 Tencent Cloud.\" />\r\n<meta name=\"MSSmartTagsPreventParsing\" content=\"True\" />\r\n<meta http-equiv=\"MSThemeCompatible\" content=\"Yes\" />\r\n<base href=\"/\" /><link rel=\"stylesheet\" type=\"text/css\" href=\"data/cache/style_1_common.css?Um1\" /><link rel=\"stylesheet\" type=\"text/css\" href=\"data/cache/style_1_forum_viewthread.css?Um1\" /><link rel=\"stylesheet\" id=\"css_extstyle\" type=\"text/css\" href=\"./template/default/style/t1/style.css\" /><script type=\"text/javascript\">var STYLEID = '1', STATICURL = 'static/', IMGDIR = 'static/image/common', VERHASH = 'Um1', charset = 'utf-8', discuz_uid = '0', cookiepre = 'cPNj_2132_', cookiedomain = '', cookiepath = '/', showusercard = '1', attackevasive = '0', disallowfloat = 'newthread', creditnotice = '3|积分|,5|金钱|,7|色币|,8|评分|', defaultstyle = './template/default/style/t1', REPORTURL = 'aHR0cHM6Ly93d3cuMThzdG0uY24vdGhyZWFkLTEyOTExMjYtMS0xLmh0bWw=', SITEURL = '/', JSPATH = 'static/js/', CSSPATH = 'data/cache/style_', DYNAMICURL = '';</script>\r\n<script src=\"static/js/common.js?Um1\" type=\"text/javascript\"></script>\r\n<script src=\"static/libs/fingerprintjs/fp.min.js\" type=\"text/javascript\" onload=\"fingerprint()\"></script>\r\n<meta name=\"application-name\" content=\"98堂[原色花堂]\" />\r\n<meta name=\"msapplication-tooltip\" content=\"98堂[原色花堂]\" />\r\n<meta name=\"msapplication-task\" content=\"name=新帖瀑布流;action-uri=/portal.php;icon-uri=static/image/common/portal.ico\" /><meta name=\"msapplication-task\" content=\"name=论坛;action-uri=/forum.php;icon-uri=static/image/common/bbs.ico\" />\r\n<link rel=\"archives\" title=\"98堂[原色花堂]\" href=\"/archiver/\" />\r\n<script src=\"static/js/forum.js?Um1\" type=\"text/javascript\"></script>\r\n</head>\r\n\r\n<body id=\"nv_forum\" class=\"pg_viewthread\" onkeydown=\"if(event.keyCode==27) return false;\">\r\n<div id=\"append_parent\"></div><div id=\"ajaxwaitid\"></div>\r\n<div id=\"toptb\" class=\"cl\">\r\n<div class=\"wp\">\r\n<div class=\"z\"><a href=\"javascript:;\"  onclick=\"setHomepage('/');\">设为首页</a><a href=\"/\"  onclick=\"addFavorite(this.href, '98堂[原色花堂]');return false;\">收藏本站</a><script src=\"static/js/language.js\" type=\"text/javascript\"></script>\r\n</div>\r\n<div class=\"y\">\r\n<a id=\"switchblind\" href=\"javascript:;\" onclick=\"toggleBlind(this)\" title=\"开启辅助访问\" class=\"switchblind\"></a>\r\n<a href=\"javascript:;\" id=\"switchwidth\" onclick=\"widthauto(this)\" title=\"切换到宽版\" class=\"switchwidth\">切换到宽版</a>\r\n</div>\r\n</div>\r\n</div>\r\n\r\n<div id=\"qmenu_menu\" class=\"p_pop blk\" style=\"display: none;\">\r\n<div class=\"ptm pbw hm\">\r\n请 <a href=\"javascript:;\" class=\"xi2\" onclick=\"lsSubmit()\"><strong>登录</strong></a> 后使用快捷导航<br />没有帐号？<a href=\"member.php?mod=register\" class=\"xi2 xw1\">立即注册</a>\r\n</div>\r\n<div id=\"fjump_menu\" class=\"btda\"></div></div><div id=\"hd\">\r\n<div class=\"wp\">\r\n<div class=\"hdc cl\"><h2><a href=\"./\" title=\"98堂[原色花堂]\"><img src=\"static/image/common/logo.png\" alt=\"98堂[原色花堂]\" border=\"0\" /></a></h2><script src=\"static/js/logging.js?Um1\" type=\"text/javascript\"></script>\r\n<form method=\"post\" autocomplete=\"off\" id=\"lsform\" action=\"/member.php?mod=logging&amp;action=login&amp;loginsubmit=yes&amp;infloat=yes&amp;lssubmit=yes\" onsubmit=\"return lsSubmit();\">\r\n<div class=\"fastlg cl\">\r\n<span id=\"return_ls\" style=\"display:none\"></span>\r\n<div class=\"y pns\">\r\n<table cellspacing=\"0\" cellpadding=\"0\">\r\n<tr>\r\n<td>\r\n<span class=\"ftid\">\r\n<select name=\"fastloginfield\" id=\"ls_fastloginfield\" width=\"40\" tabindex=\"900\">\r\n<option value=\"username\">用户名</option>\r\n<option value=\"email\">Email</option>\r\n</select>\r\n</span>\r\n<script type=\"text/javascript\">simulateSelect('ls_fastloginfield')</script>\r\n</td>\r\n<td><input type=\"text\" name=\"username\" id=\"ls_username\" autocomplete=\"off\" class=\"px vm\" tabindex=\"901\" /></td>\r\n<td class=\"fastlg_l\"><label for=\"ls_cookietime\"><input type=\"checkbox\" name=\"cookietime\" id=\"ls_cookietime\" class=\"pc\" value=\"2592000\" tabindex=\"903\" checked=\"checked\"/>自动登录</label></td>\r\n<td>&nbsp;<a href=\"javascript:;\" onclick=\"showWindow('login', 'member.php?mod=logging&action=login&viewlostpw=1')\">找回密码</a></td>\r\n</tr>\r\n<tr>\r\n<td><label for=\"ls_password\" class=\"z psw_w\">密码</label></td>\r\n<td><input type=\"password\" name=\"password\" id=\"ls_password\" class=\"px vm\" autocomplete=\"off\" tabindex=\"902\" /></td>\r\n<td class=\"fastlg_l\"><button type=\"submit\" class=\"pn vm\" tabindex=\"904\" style=\"width: 75px;\"><em>登录</em></button></td>\r\n<td>&nbsp;<a href=\"member.php?mod=register\" class=\"xi2 xw1\">立即注册</a></td>\r\n</tr>\r\n</table>\r\n<input type=\"hidden\" name=\"formhash\" value=\"786d433d\" />\r\n<input type=\"hidden\" name=\"quickforward\" value=\"yes\" />\r\n<input type=\"hidden\" name=\"handlekey\" value=\"ls\" />\r\n</div>\r\n</div>\r\n</form>\r\n\r\n</div>\r\n\r\n<div id=\"nv\">\r\n<a href=\"javascript:;\" id=\"qmenu\" onmouseover=\"delayShow(this, function () {showMenu({'ctrlid':'qmenu','pos':'34!','ctrlclass':'a','duration':2});showForummenu(36);})\">快捷导航</a>\r\n<ul><li class=\"a\" id=\"mn_forum\" ><a href=\"forum.php\" hidefocus=\"true\" title=\"BBS\"  >论坛<span>BBS</span></a></li><li id=\"mn_N743b\" ><a href=\"book.html\" hidefocus=\"true\"  >小说</a></li><li id=\"mn_portal\" ><a href=\"portal.php\" hidefocus=\"true\" title=\"Portal\"  >新帖瀑布流<span>Portal</span></a></li></ul>\r\n</div>\r\n<div class=\"p_pop h_pop\" id=\"mn_userapp_menu\" style=\"display: none\"></div><div id=\"mu\" class=\"cl\">\r\n</div><div id=\"scbar\" class=\"cl\">\r\n<form id=\"scbar_form\" method=\"post\" autocomplete=\"off\" onsubmit=\"searchFocus($('scbar_txt'))\" action=\"search.php?searchsubmit=yes\" target=\"_blank\">\r\n<input type=\"hidden\" name=\"mod\" id=\"scbar_mod\" value=\"search\" />\r\n<input type=\"hidden\" name=\"formhash\" value=\"786d433d\" />\r\n<input type=\"hidden\" name=\"srchtype\" value=\"title\" />\r\n<input type=\"hidden\" name=\"srhfid\" value=\"36\" />\r\n<input type=\"hidden\" name=\"srhlocality\" value=\"forum::viewthread\" />\r\n<table cellspacing=\"0\" cellpadding=\"0\">\r\n<tr>\r\n<td class=\"scbar_icon_td\"></td>\r\n<td class=\"scbar_txt_td\"><input type=\"text\" name=\"srchtxt\" id=\"scbar_txt\" value=\"请输入搜索内容\" autocomplete=\"off\" x-webkit-speech speech /></td>\r\n<td class=\"scbar_type_td\"><a href=\"javascript:;\" id=\"scbar_type\" class=\"xg1\" onclick=\"showMenu(this.id)\" hidefocus=\"true\">搜索</a></td>\r\n<td class=\"scbar_btn_td\"><button type=\"submit\" name=\"searchsubmit\" id=\"scbar_btn\" sc=\"1\" class=\"pn pnc\" value=\"true\"><strong class=\"xi2\">搜索</strong></button></td>\r\n<td class=\"scbar_hot_td\">\r\n<div id=\"scbar_hot\">\r\n</div>\r\n</td>\r\n</tr>\r\n</table>\r\n</form>\r\n</div>\r\n<ul id=\"scbar_type_menu\" class=\"p_pop\" style=\"display: none;\"><li><a href=\"javascript:;\" rel=\"curforum\" fid=\"36\" >本版</a></li><li><a href=\"javascript:;\" rel=\"forum\" class=\"curtype\">帖子</a></li><li><a href=\"javascript:;\" rel=\"user\">用户</a></li></ul>\r\n<script type=\"text/javascript\">\r\ninitSearchmenu('scbar', '');\r\n</script>\r\n</div>\r\n</div>\r\n\r\n\r\n<div id=\"wp\" class=\"wp\">\r\n<div class=\"h10\"></div>\r\n<div class=\"show-text cl\">\r\n        <a href=\"https://tz112.com\" target=\"_blank\" class=\"js-randomBg\">亚博美女赌场</a>\r\n        <a href=\"https://5656456565.com/vote_topic_80615.do\" target=\"_blank\" class=\"js-randomBg\">澳门现金赌场</a>\r\n        <a href=\"http://23.225.52.51:4466/vip103.html\" target=\"_blank\" class=\"js-randomBg\">澳门新葡京</a>\r\n        <a href=\"https://by301.hdzshbk.cn/?channelCode=by_147\" target=\"_blank\" class=\"js-randomBg\">鲍鱼全球黄播</a>\r\n        <a href=\"http://23.224.188.2:2939/vip107.html\" target=\"_blank\" class=\"js-randomBg\">澳门威尼斯人</a>\r\n        <a href=\"https://154.82.93.23:443/vip058\" target=\"_blank\" class=\"js-randomBg\">澳门葡京赌场</a>\r\n        <a href=\"https://66861056.app/registered\" target=\"_blank\" class=\"js-randomBg\">天天送红包</a>\r\n        <a href=\"https://5581268.cc:8443/?shareName=5581268.cc\" target=\"_blank\" class=\"js-randomBg\">【开元棋牌】</a>\r\n        <a href=\"https://h1116.cc:15555\" target=\"_blank\" class=\"js-randomBg\"><font color=\"#00ff00\">现金棋牌</font></a>\r\n        <a href=\"https://www.rckif.com/001.html\" target=\"_blank\" class=\"js-randomBg\">91成人抖音</a>\r\n    </div>\r\n<script type=\"text/javascript\">var fid = parseInt('36'), tid = parseInt('1291126');</script>\r\n\r\n<script src=\"static/js/forum_viewthread.js?Um1\" type=\"text/javascript\"></script>\r\n<script type=\"text/javascript\">zoomstatus = parseInt(1);var imagemaxwidth = '600';var aimgcount = new Array();</script>\r\n\r\n<style id=\"diy_style\" type=\"text/css\"></style>\r\n<!--[diy=diynavtop]--><div id=\"diynavtop\" class=\"area\"></div><!--[/diy]-->\r\n<div id=\"pt\" class=\"bm cl\">\r\n<div class=\"z\">\r\n<a href=\"./\" class=\"nvhm\" title=\"首页\">98堂[原色花堂]</a><em>&raquo;</em><a href=\"forum.php\">论坛</a> <em>&rsaquo;</em> <a href=\"forum.php?gid=1\">原创BT电影</a> <em>&rsaquo;</em> <a href=\"forum-36-1.html\">亚洲无码原创</a> <em>&rsaquo;</em> <a href=\"thread-1291126-1-1.html\">SSNI-116  [无码破解] 絶頂してピクピクしているおま● ...</a>\r\n</div>\r\n</div>\r\n\r\n<style id=\"diy_style\" type=\"text/css\"></style>\r\n<div class=\"wp\">\r\n<!--[diy=diy1]--><div id=\"diy1\" class=\"area\"></div><!--[/diy]-->\r\n</div>\r\n\r\n<div id=\"ct\" class=\"wp cl\">\r\n<div id=\"pgt\" class=\"pgs mbm cl \">\r\n<div class=\"pgt\"></div>\r\n<span class=\"y pgb\" id=\"visitedforums\" onmouseover=\"$('visitedforums').id = 'visitedforumstmp';this.id = 'visitedforums';showMenu({'ctrlid':this.id,'pos':'34'})\"><a href=\"forum-36-1.html\">返回列表</a></span>\r\n<a id=\"newspecial\" onmouseover=\"$('newspecial').id = 'newspecialtmp';this.id = 'newspecial';showMenu({'ctrlid':this.id})\" onclick=\"showWindow('newthread', 'forum.php?mod=post&action=newthread&fid=36')\" href=\"javascript:;\" title=\"发新帖\"><img src=\"static/image/common/pn_post.png\" alt=\"发新帖\" /></a></div>\r\n\r\n\r\n\r\n<div id=\"postlist\" class=\"pl bm\">\r\n<table cellspacing=\"0\" cellpadding=\"0\">\r\n<tr>\r\n<td class=\"pls ptn pbn\">\r\n<div class=\"hm ptn\">\r\n<span class=\"xg1\">查看:</span> <span class=\"xi1\">10248</span><span class=\"pipe\">|</span><span class=\"xg1\">回复:</span> <span class=\"xi1\">3</span>\r\n</div>\r\n</td>\r\n<td class=\"plc ptm pbn vwthd\">\r\n<div class=\"y\">\r\n<a href=\"forum.php?mod=viewthread&amp;action=printable&amp;tid=1291126\" title=\"打印\" target=\"_blank\"><img src=\"static/image/common/print.png\" alt=\"打印\" class=\"vm\" /></a>\r\n<a href=\"forum.php?mod=redirect&amp;goto=nextoldset&amp;tid=1291126\" title=\"上一主题\"><img src=\"static/image/common/thread-prev.png\" alt=\"上一主题\" class=\"vm\" /></a>\r\n<a href=\"forum.php?mod=redirect&amp;goto=nextnewset&amp;tid=1291126\" title=\"下一主题\"><img src=\"static/image/common/thread-next.png\" alt=\"下一主题\" class=\"vm\" /></a>\r\n</div>\r\n<h1 class=\"ts\">\r\n<a href=\"forum.php?mod=forumdisplay&amp;fid=36&amp;filter=typeid&amp;typeid=672\">[无码破解]</a>\r\n<span id=\"thread_subject\">SSNI-116  [无码破解] 絶頂してピクピクしているおま●こを容赦なく突きまくる怒涛の</span>\r\n</h1>\r\n<span class=\"xg1\">\r\n<a href=\"thread-1291126-1-1.html\" onclick=\"return copyThreadUrl(this, '98堂[原色花堂]')\" >[复制链接]</a>\r\n</span>\r\n</td>\r\n</tr>\r\n</table>\r\n\r\n\r\n<table cellspacing=\"0\" cellpadding=\"0\" class=\"ad\">\r\n<tr>\r\n<td class=\"pls\">\r\n</td>\r\n<td class=\"plc\">\r\n</td>\r\n</tr>\r\n</table><style>\r\n.show-text4{position:relative;padding-left:170px;min-height:52px}\r\n.show-text4 .show-info{position:absolute;top:0;left:0;width:160px;height:100%;text-align:center;display:block;border-right:1px solid #c2d5e3}\r\n.show-text4 .show-info .avatar{position:absolute;top:50%;left:50%;margin-top:-23px;margin-left:-23px;width:46px;height:46px;line-height:46px;-webkit-border-radius:5rem;border-radius:5rem;background:#c1c1c1;color:#fff;text-align:center;font-size:18px}\r\n.show-text4 .show-main{padding:5px 10px}\r\n.show-text4 .item{width:50%;float:left;display:block;text-decoration:none}\r\n.show-text4 .item h4{font-size:16px;color:#36f}\r\n.show-text4 .item p{color:#777}\r\n</style><div id=\"post_10775860\" ><table id=\"pid10775860\" class=\"plhin\" summary=\"pid10775860\" cellspacing=\"0\" cellpadding=\"0\">\r\n<tr>\r\n<a name=\"newpost\"></a> <td class=\"pls\" rowspan=\"2\">\r\n<div id=\"favatar10775860\" class=\"pls cl favatar\">\r\n<div class=\"pi\">\r\n<div class=\"authi\"><a href=\"space-uid-420458.html\" target=\"_blank\" class=\"xw1\" style=\"color: #0033FF\">junmin1995</a>\r\n</div>\r\n</div>\r\n<div class=\"p_pop blk bui card_gender_0\" id=\"userinfo10775860\" style=\"display: none; margin-top: -11px;\">\r\n<div class=\"m z\">\r\n<div id=\"userinfo10775860_ma\"></div>\r\n</div>\r\n<div class=\"i y\">\r\n<div>\r\n<strong><a href=\"space-uid-420458.html\" target=\"_blank\" class=\"xi2\" style=\"color: #0033FF\">junmin1995</a></strong>\r\n<em>当前在线</em>\r\n</div><dl class=\"cl\">\r\n<dt>积分</dt><dd><a href=\"home.php?mod=space&uid=420458&do=profile\" target=\"_blank\" class=\"xi2\">63788</a></dd>\r\n</dl><div class=\"imicn\">\r\n<a href=\"home.php?mod=space&amp;uid=420458&amp;do=profile\" target=\"_blank\" title=\"查看详细资料\"><img src=\"static/image/common/userinfo.gif\" alt=\"查看详细资料\" /></a>\r\n</div>\r\n<div id=\"avatarfeed\"><span id=\"threadsortswait\"></span></div>\r\n</div>\r\n</div>\r\n<div>\r\n<div class=\"avatar\" onmouseover=\"showauthor(this, 'userinfo10775860')\"><a href=\"space-uid-420458.html\" class=\"avtm\" target=\"_blank\"><img src=\"/uc_server/data/avatar/000/42/04/58_avatar_middle.jpg\" onerror=\"this.onerror=null;this.src='/uc_server/images/noavatar_middle.gif'\" /></a></div>\r\n</div>\r\n<div class=\"tns xg2\"><table cellspacing=\"0\" cellpadding=\"0\"><th><p><a href=\"home.php?mod=space&uid=420458&do=profile\" class=\"xi2\"><span title=\"64221\">6万</span></a></p>金钱</th><th><p><a href=\"home.php?mod=space&uid=420458&do=profile\" class=\"xi2\">0</a></p>色币</th><td><p><a href=\"home.php?mod=space&uid=420458&do=profile\" class=\"xi2\"><span title=\"213588\">21万</span></a></p>评分</td></table></div>\r\n\r\n<p><em><a href=\"home.php?mod=spacecp&amp;ac=usergroup&amp;gid=3\" target=\"_blank\"><font color=\"#0033FF\">版主</font></a></em></p>\r\n\r\n\r\n<dl class=\"pil cl\">\r\n  <dt>积分</dt><dd><a href=\"home.php?mod=space&uid=420458&do=profile\" target=\"_blank\" class=\"xi2\">63788</a></dd>\r\n</dl>\r\n\r\n\r\n<ul class=\"xl xl2 o cl\">\r\n<li class=\"pm2\"><a href=\"home.php?mod=spacecp&amp;ac=pm&amp;op=showmsg&amp;handlekey=showmsg_420458&amp;touid=420458&amp;pmid=0&amp;daterange=2&amp;pid=10775860&amp;tid=1291126\" onclick=\"showWindow('sendpm', this.href);\" title=\"发消息\" class=\"xi2\">发消息</a></li>\r\n</ul>\r\n</div>\r\n</td>\r\n<td class=\"plc\">\r\n<div class=\"pi\">\r\n<div id=\"fj\" class=\"y\">\r\n<label class=\"z\">电梯直达</label>\r\n<input type=\"text\" class=\"px p_fre z\" size=\"2\" onkeyup=\"$('fj_btn').href='forum.php?mod=redirect&ptid=1291126&authorid=0&postno='+this.value\" onkeydown=\"if(event.keyCode==13) {window.location=$('fj_btn').href;return false;}\" title=\"跳转到指定楼层\" />\r\n<a href=\"javascript:;\" id=\"fj_btn\" class=\"z\" title=\"跳转到指定楼层\"><img src=\"static/image/common/fj_btn.png\" alt=\"跳转到指定楼层\" class=\"vm\" /></a>\r\n</div>\r\n<strong>\r\n<a href=\"thread-1291126-1-1.html\"   id=\"postnum10775860\" onclick=\"setCopy(this.href, '帖子地址复制成功');return false;\">\r\n楼主</a>\r\n</strong>\r\n<div class=\"pti\">\r\n<div class=\"pdbt\">\r\n</div>\r\n<div class=\"authi\">\r\n<img class=\"authicn vm\" id=\"authicon10775860\" src=\"static/image/common/online_moderator.gif\" />\r\n<em id=\"authorposton10775860\">发表于 <span title=\"2023-04-28 09:02:02\">昨天&nbsp;09:02</span></em>\r\n<span class=\"pipe\">|</span>\r\n<a href=\"forum.php?mod=viewthread&amp;tid=1291126&amp;page=1&amp;authorid=420458\" rel=\"nofollow\">只看该作者</a>\r\n<span class=\"pipe\">|</span><a href=\"forum.php?mod=viewthread&amp;tid=1291126&amp;from=album\">只看大图</a>\r\n<span class=\"none\"><img src=\"static/image/common/arw_r.gif\" class=\"vm\" alt=\"回帖奖励\" /></span>\r\n<span class=\"pipe show\">|</span><a href=\"forum.php?mod=viewthread&amp;tid=1291126&amp;extra=page%3D1&amp;ordertype=1\"  class=\"show\">倒序浏览</a>\r\n<span class=\"pipe show\">|</span><a href=\"javascript:;\" onclick=\"readmode($('thread_subject').innerHTML, 10775860);\" class=\"show\">阅读模式</a>\r\n</div>\r\n</div>\r\n</div><div class=\"pct\"><style type=\"text/css\">.pcb{margin-right:0}</style><div class=\"pcb\">\r\n<div class=\"t_fsz\">\r\n<table cellspacing=\"0\" cellpadding=\"0\"><tr><td class=\"t_f\" id=\"postmessage_10775860\">\r\n【影片名称】：絶頂してピクピクしているおま●こを容赦なく突きまくる怒涛のおかわり激ピストン性交 <br />\r\n【出演女优】：天使もえ<br />\r\n【影片容量】：5.07G<br />\r\n【是否有码】：无码<br />\r\n【种子期限】：5种或健康度1000<br />\r\n【下载工具】：比特彗星 比特精灵 uTorrent QBittorrent 迅雷极速版 闪电下载【请不要用迅雷官方版下载，官方版本已经被屏蔽】<br />\r\n【影片预览】：看不到图请挂代理或点右键显示图片<br />\r\n\r\n<ignore_js_op>\r\n\r\n<img aid=\"4694219\" src=\"static/image/common/none.gif\" zoomfile=\"https://xksm54s.com/tupian/forum/202304/28/090142n7uumun66abf6f6a.jpg\" file=\"https://xksm54s.com/tupian/forum/202304/28/090142n7uumun66abf6f6a.jpg\" class=\"zoom\" onclick=\"zoom(this, this.src, 0, 0, 0)\" width=\"600\" id=\"aimg_4694219\" inpost=\"1\" onmouseover=\"showMenu({'ctrlid':this.id,'pos':'12'})\" />\r\n\r\n<div class=\"tip tip_4 aimg_tip\" id=\"aimg_4694219_menu\" style=\"position: absolute; display: none\" disautofocus=\"true\">\r\n<div class=\"xs0\">\r\n<p><strong>SSNI-116.jpg</strong> <em class=\"xg1\">(178.25 KB, 下载次数: 0)</em></p>\r\n<p>\r\n\r\n<a href=\"https://xksm54s.com/tupian/forum/202304/28/090142n7uumun66abf6f6a.jpg\" target=\"_blank\">下载附件</a>\r\n\r\n</p>\r\n\r\n<p class=\"xg1 y\"><span title=\"2023-04-28 09:01\">昨天&nbsp;09:01</span> 上传</p>\r\n\r\n</div>\r\n<div class=\"tip_horn\"></div>\r\n</div>\r\n\r\n</ignore_js_op>\r\n<br />\r\n\r\n<ignore_js_op>\r\n\r\n<img aid=\"4694220\" src=\"static/image/common/none.gif\" zoomfile=\"https://xksm54s.com/tupian/forum/202304/28/090142z3b3abk35ogdqhon.jpg\" file=\"https://xksm54s.com/tupian/forum/202304/28/090142z3b3abk35ogdqhon.jpg\" class=\"zoom\" onclick=\"zoom(this, this.src, 0, 0, 0)\" width=\"600\" id=\"aimg_4694220\" inpost=\"1\" onmouseover=\"showMenu({'ctrlid':this.id,'pos':'12'})\" />\r\n\r\n<div class=\"tip tip_4 aimg_tip\" id=\"aimg_4694220_menu\" style=\"position: absolute; display: none\" disautofocus=\"true\">\r\n<div class=\"xs0\">\r\n<p><strong>SSNI-116-U.jpg</strong> <em class=\"xg1\">(315.05 KB, 下载次数: 0)</em></p>\r\n<p>\r\n\r\n<a href=\"https://xksm54s.com/tupian/forum/202304/28/090142z3b3abk35ogdqhon.jpg\" target=\"_blank\">下载附件</a>\r\n\r\n</p>\r\n\r\n<p class=\"xg1 y\"><span title=\"2023-04-28 09:01\">昨天&nbsp;09:01</span> 上传</p>\r\n\r\n</div>\r\n<div class=\"tip_horn\"></div>\r\n</div>\r\n\r\n</ignore_js_op>\r\n<br />\r\n磁力链接：<br />\r<div class=\"blockcode\"><div id=\"code_aV3\"><ol><li>magnet:?xt=urn:btih:6D7547B708F10780DF910B081168A2BE831D859F</ol></div><em onclick=\"copycode($('code_aV3'));\">复制代码</em></div><br />\n\r<br />\n\r<br />\n</td></tr></table>\r\n\r\n<div class=\"modact\"><a href=\"forum.php?mod=misc&amp;action=viewthreadmod&amp;tid=1291126\" title=\"帖子模式\" onclick=\"showWindow('viewthreadmod', this.href)\">本主题由 junmin1995 于 <span title=\"2023-04-28 09:23\">昨天&nbsp;09:23</span> 限时高亮</a></div><div class=\"pattl\">\r\n\r\n<ignore_js_op>\r\n<dl class=\"tattl\">\r\n<dt>\r\n<img src=\"static/image/filetype/torrent.gif\" border=\"0\" class=\"vm\" alt=\"\" />\r\n</dt>\r\n<dd>\r\n<p class=\"attnm\">\r\n\r\n<a href=\"https://xksm54s.com/tupian/down_v2.php?sign=eyJ2YWx1ZSI6Im03SE9YUXJQRHJaK0pFVlRGemNjSG56SkZPUkgyTDZ0MlhiVXJGcG05UFA2cDdETG5tdkkrQ0hwR0RqbmZFenRkVTlSR0hnNk5IelJqZm9lVVBuelBJcCtcL1B3OUZNTU5ucklNazdURFJHR21hWEltK3V6M1lWWHBLRGdsZ0xVYzFsemVMaktvZzFnc0pTNHhYYXlteVdFWWhBZ1FHZm1YQmJMNFVzOGc2QjJsMkJsV2VLKzVMUE9lMVZzVCtcL3FyUW1Sajg2WEhsdVN0QThJcUxzWDhvRFdSUUN0RXNQeUYxUisxQXdlbk5McDl2MUE3U1ZIOUVzWnRVUkRoMnpsWSIsIml2IjoiSG15WjgyRlZ0QzlTS3FVNCJ9\" onmouseover=\"showMenu({'ctrlid':this.id,'pos':'12'})\" id=\"aid4694221\" target=\"_blank\">SSNI-116-U.torrent</a>\r\n\r\n<div class=\"tip tip_4\" id=\"aid4694221_menu\" style=\"display: none\" disautofocus=\"true\">\r\n<div class=\"tip_c\">\r\n<p class=\"y\"><span title=\"2023-04-28 09:01\">昨天&nbsp;09:01</span> 上传</p>\r\n<p>点击文件名下载附件</p>\r\n\r\n</div>\r\n<div class=\"tip_horn\"></div>\r\n</div>\r\n</p>\r\n<p>27.78 KB</p>\r\n<p>\r\n\r\n</p>\r\n\r\n\r\n</dd>\r\n</dl>\r\n</ignore_js_op>\r\n</div>\r\n\r\n</div>\r\n<div id=\"comment_10775860\" class=\"cm\">\r\n</div>\r\n\r\n<h3 class=\"psth xs1\"><span class=\"icon_ring vm\"></span>评分</h3>\r\n<dl id=\"ratelog_10775860\" class=\"rate\">\r\n<dd style=\"margin:0\">\r\n<div id=\"post_rate_10775860\"></div>\r\n<table class=\"ratl\">\r\n<tr>\r\n<th class=\"xw1\" width=\"120\"><a href=\"forum.php?mod=misc&amp;action=viewratings&amp;tid=1291126&amp;pid=10775860\" onclick=\"showWindow('viewratings', this.href)\" title=\"查看全部评分\"> 参与人数 <span class=\"xi1\">5</span></a></th><th class=\"xw1\" width=\"80\">评分 <i><span class=\"xi1\">+12</span></i></th>\r\n<th>\r\n<a href=\"javascript:;\" onclick=\"toggleRatelogCollapse('ratelog_10775860', this);\" class=\"y xi2 op\">收起</a>\r\n<i class=\"txt_h\">理由</i>\r\n</th>\r\n</tr>\r\n<tbody class=\"ratl_l\"><tr id=\"rate_10775860_435606\">\r\n<td>\r\n<a href=\"space-uid-435606.html\" target=\"_blank\"><img src=\"/uc_server/data/avatar/000/43/56/06_avatar_small.jpg\" onerror=\"this.onerror=null;this.src='/uc_server/images/noavatar_small.gif'\" /></a> <a href=\"space-uid-435606.html\" target=\"_blank\">mojo</a>\r\n</td><td class=\"xi1\"> + 3</td>\r\n<td class=\"xg1\"></td>\r\n</tr>\r\n<tr id=\"rate_10775860_447195\">\r\n<td>\r\n<a href=\"space-uid-447195.html\" target=\"_blank\"><img src=\"/uc_server/data/avatar/000/44/71/95_avatar_small.jpg\" onerror=\"this.onerror=null;this.src='/uc_server/images/noavatar_small.gif'\" /></a> <a href=\"space-uid-447195.html\" target=\"_blank\">yoscso</a>\r\n</td><td class=\"xi1\"> + 3</td>\r\n<td class=\"xg1\">很给力!</td>\r\n</tr>\r\n<tr id=\"rate_10775860_431561\">\r\n<td>\r\n<a href=\"space-uid-431561.html\" target=\"_blank\"><img src=\"/uc_server/data/avatar/000/43/15/61_avatar_small.jpg\" onerror=\"this.onerror=null;this.src='/uc_server/images/noavatar_small.gif'\" /></a> <a href=\"space-uid-431561.html\" target=\"_blank\">su133</a>\r\n</td><td class=\"xi1\"> + 3</td>\r\n<td class=\"xg1\">很给力!</td>\r\n</tr>\r\n<tr id=\"rate_10775860_433632\">\r\n<td>\r\n<a href=\"space-uid-433632.html\" target=\"_blank\"><img src=\"/uc_server/data/avatar/000/43/36/32_avatar_small.jpg\" onerror=\"this.onerror=null;this.src='/uc_server/images/noavatar_small.gif'\" /></a> <a href=\"space-uid-433632.html\" target=\"_blank\">色色色鳄鱼</a>\r\n</td><td class=\"xi1\"> + 2</td>\r\n<td class=\"xg1\">很给力!</td>\r\n</tr>\r\n<tr id=\"rate_10775860_433163\">\r\n<td>\r\n<a href=\"space-uid-433163.html\" target=\"_blank\"><img src=\"/uc_server/data/avatar/000/43/31/63_avatar_small.jpg\" onerror=\"this.onerror=null;this.src='/uc_server/images/noavatar_small.gif'\" /></a> <a href=\"space-uid-433163.html\" target=\"_blank\">123cc</a>\r\n</td><td class=\"xi1\"> + 1</td>\r\n<td class=\"xg1\">很给力!</td>\r\n</tr>\r\n</tbody>\r\n</table>\r\n<p class=\"ratc\">\r\n<a href=\"forum.php?mod=misc&amp;action=viewratings&amp;tid=1291126&amp;pid=10775860\" onclick=\"showWindow('viewratings', this.href)\" title=\"查看全部评分\" class=\"xi2\">查看全部评分</a>\r\n</p>\r\n</dd>\r\n</dl>\r\n</div>\r\n</div>\r\n\r\n</td></tr>\r\n<tr><td class=\"plc plm\">\r\n\r\n<div id=\"p_btn\" class=\"mtw mbm hm cl\">\r\n\r\n<a href=\"home.php?mod=spacecp&amp;ac=favorite&amp;type=thread&amp;id=1291126&amp;formhash=786d433d\" id=\"k_favorite\" onclick=\"showWindow(this.id, this.href, 'get', 0);\" onmouseover=\"this.title = $('favoritenumber').innerHTML + ' 人收藏'\" title=\"收藏本帖\"><i><img src=\"static/image/common/fav.gif\" alt=\"收藏\" />收藏<span id=\"favoritenumber\">6</span></i></a>\r\n</div>\r\n\r\n</td>\r\n</tr>\r\n<tr id=\"_postposition10775860\"></tr>\r\n<tr>\r\n<td class=\"pls\"></td>\r\n<td class=\"plc\" style=\"overflow:visible;\">\r\n<div class=\"po hin\">\r\n<div class=\"pob cl\">\r\n<em>\r\n<a class=\"fastre\" href=\"forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1291126&amp;reppost=10775860&amp;extra=page%3D1&amp;page=1\" onclick=\"showWindow('reply', this.href)\">回复</a>\r\n</em>\r\n\r\n<p>\r\n<a href=\"javascript:;\" id=\"mgc_post_10775860\" onmouseover=\"showMenu(this.id)\" class=\"showmenu\">使用道具</a>\r\n<a href=\"javascript:;\" onclick=\"showWindow('miscreport10775860', 'misc.php?mod=report&rtype=post&rid=10775860&tid=1291126&fid=36', 'get', -1);return false;\">举报</a>\r\n</p>\r\n\r\n<ul id=\"mgc_post_10775860_menu\" class=\"p_pop mgcmn\" style=\"display: none;\">\r\n</ul>\r\n<script type=\"text/javascript\" reload=\"1\">checkmgcmn('post_10775860')</script>\r\n</div>\r\n</div>\r\n</td>\r\n</tr>\r\n<tr class=\"ad\">\r\n<td class=\"pls\">\r\n</td>\r\n<td class=\"plc\">\r\n</td>\r\n</tr>\r\n</table>\r\n<script type=\"text/javascript\" reload=\"1\">\r\naimgcount[10775860] = ['4694219','4694220'];\r\nattachimggroup(10775860);\r\nvar aimgfid = 0;\r\n</script>\r\n</div>\r\n\r\n<div class=\"show-text2 pad-tb-10\"><a href=\"https://tz112.com\" target=\"_blank\">亚博美女赌场</a>\r\n<a href=\"https://5656456565.com/vote_topic_80615.do\" target=\"_blank\">澳门现金赌场</a>\r\n<a href=\"http://23.225.52.51:4466/vip103.html\" target=\"_blank\">澳门新葡京</a>\r\n<a href=\"https://by301.hdzshbk.cn/?channelCode=by_147\" target=\"_blank\">鲍鱼全球黄播</a>\r\n<a href=\"http://23.224.188.2:2939/vip107.html\" target=\"_blank\">澳门威尼斯人</a>\r\n<a href=\"https://154.82.93.23:443/vip058\" target=\"_blank\">澳门葡京赌场</a>\r\n<a href=\"https://66861056.app/registered\" target=\"_blank\">天天送红包</a>\r\n<a href=\"https://5581268.cc:8443/?shareName=5581268.cc\" target=\"_blank\">【开元棋牌】</a>\r\n<a href=\"https://h1116.cc:15555\" target=\"_blank\">现金棋牌</a>\r\n<a href=\"https://www.rckif.com/001.html\" target=\"_blank\">91成人抖音</a>\r\n</div>\r\n<table class=\"plhin\">\r\n<tr class=\"ad\">\r\n<td class=\"pls\"></td>\r\n<td class=\"plc\"></td>\r\n</tr>\r\n</table>\r\n<div class=\"show-text4 pad-tb-10\">\r\n<div class=\"show-info\">\r\n<div class=\"avatar\">AD</div>\r\n</div>\r\n<div class=\"show-main\"><div><a href=\"https://www.d23dx.com\" class=\"item js-appJump\">\r\n    <h4>鲍鱼盒子，全网高清直播资源‖自拍视频在线看</h4>\r\n    <p>注册免费观看，邀请好友送现金红包。</p>\r\n</a>\r\n<a href=\"https://www.d23dx.com\" class=\"item js-appJump\">\r\n    <h4>鲍鱼盒子，全网高清直播资源‖自拍视频在线看</h4>\r\n    <p>注册免费观看，邀请好友送现金红包。</p>\r\n</a></div></div>\r\n</div>\r\n<table class=\"plhin\">\r\n<tr class=\"ad\">\r\n<td class=\"pls\"></td>\r\n<td class=\"plc\"></td>\r\n</tr>\r\n</table>\r\n<div id=\"post_10776497\" ><table id=\"pid10776497\" class=\"plhin\" summary=\"pid10776497\" cellspacing=\"0\" cellpadding=\"0\">\r\n<tr>\r\n <td class=\"pls\" rowspan=\"2\">\r\n<div id=\"favatar10776497\" class=\"pls cl favatar\">\r\n<div class=\"pi\">\r\n<div class=\"authi\"><a href=\"space-uid-445078.html\" target=\"_blank\" class=\"xw1\">张闹闹</a>\r\n</div>\r\n</div>\r\n<div class=\"p_pop blk bui card_gender_0\" id=\"userinfo10776497\" style=\"display: none; margin-top: -11px;\">\r\n<div class=\"m z\">\r\n<div id=\"userinfo10776497_ma\"></div>\r\n</div>\r\n<div class=\"i y\">\r\n<div>\r\n<strong><a href=\"space-uid-445078.html\" target=\"_blank\" class=\"xi2\">张闹闹</a></strong>\r\n<em>当前在线</em>\r\n</div><dl class=\"cl\">\r\n<dt>积分</dt><dd><a href=\"home.php?mod=space&uid=445078&do=profile\" target=\"_blank\" class=\"xi2\">25</a></dd>\r\n</dl><div class=\"imicn\">\r\n<a href=\"home.php?mod=space&amp;uid=445078&amp;do=profile\" target=\"_blank\" title=\"查看详细资料\"><img src=\"static/image/common/userinfo.gif\" alt=\"查看详细资料\" /></a>\r\n</div>\r\n<div id=\"avatarfeed\"><span id=\"threadsortswait\"></span></div>\r\n</div>\r\n</div>\r\n<div>\r\n<div class=\"avatar\" onmouseover=\"showauthor(this, 'userinfo10776497')\"><a href=\"space-uid-445078.html\" class=\"avtm\" target=\"_blank\"><img src=\"/uc_server/data/avatar/000/44/50/78_avatar_middle.jpg\" onerror=\"this.onerror=null;this.src='/uc_server/images/noavatar_middle.gif'\" /></a></div>\r\n</div>\r\n<div class=\"tns xg2\"><table cellspacing=\"0\" cellpadding=\"0\"><th><p><a href=\"home.php?mod=space&uid=445078&do=profile\" class=\"xi2\">94</a></p>金钱</th><th><p><a href=\"home.php?mod=space&uid=445078&do=profile\" class=\"xi2\">0</a></p>色币</th><td><p><a href=\"home.php?mod=space&uid=445078&do=profile\" class=\"xi2\">0</a></p>评分</td></table></div>\r\n\r\n<p><em><a href=\"home.php?mod=spacecp&amp;ac=usergroup&amp;gid=11\" target=\"_blank\">Lv2 无名之辈</a></em></p>\r\n\r\n\r\n<p><span class=\"pbg2\"  id=\"upgradeprogress_10776497\" onmouseover=\"showMenu({'ctrlid':this.id, 'pos':'12!', 'menuid':'g_up10776497_menu'});\"><span class=\"pbr2\" style=\"width:75%;\"></span></span></p>\r\n<div id=\"g_up10776497_menu\" class=\"tip tip_4\" style=\"display: none;\"><div class=\"tip_horn\"></div><div class=\"tip_c\">Lv2 无名之辈, 积分 25, 距离下一级还需 5 积分</div></div>\r\n\r\n<dl class=\"pil cl\">\r\n  <dt>积分</dt><dd><a href=\"home.php?mod=space&uid=445078&do=profile\" target=\"_blank\" class=\"xi2\">25</a></dd>\r\n</dl>\r\n\r\n\r\n<ul class=\"xl xl2 o cl\">\r\n<li class=\"pm2\"><a href=\"home.php?mod=spacecp&amp;ac=pm&amp;op=showmsg&amp;handlekey=showmsg_445078&amp;touid=445078&amp;pmid=0&amp;daterange=2&amp;pid=10776497&amp;tid=1291126\" onclick=\"showWindow('sendpm', this.href);\" title=\"发消息\" class=\"xi2\">发消息</a></li>\r\n</ul>\r\n</div>\r\n</td>\r\n<td class=\"plc\">\r\n<div class=\"pi\">\r\n<strong>\r\n<a href=\"forum.php?mod=redirect&goto=findpost&ptid=1291126&pid=10776497\"   id=\"postnum10776497\" onclick=\"setCopy(this.href, '帖子地址复制成功');return false;\">\r\n沙发</a>\r\n</strong>\r\n<div class=\"pti\">\r\n<div class=\"pdbt\">\r\n</div>\r\n<div class=\"authi\">\r\n<img class=\"authicn vm\" id=\"authicon10776497\" src=\"static/image/common/online_member.gif\" />\r\n<em id=\"authorposton10776497\">发表于 <span title=\"2023-04-28 09:59:17\">昨天&nbsp;09:59</span></em>\r\n<span class=\"xg1\">来自手机</span>\r\n<span class=\"pipe\">|</span>\r\n<a href=\"forum.php?mod=viewthread&amp;tid=1291126&amp;page=1&amp;authorid=445078\" rel=\"nofollow\">只看该作者</a>\r\n</div>\r\n</div>\r\n</div><div class=\"pct\"><div class=\"pcb\">\r\n<div class=\"t_fsz\">\r\n<table cellspacing=\"0\" cellpadding=\"0\"><tr><td class=\"t_f\" id=\"postmessage_10776497\">\r\n这就很顶了</td></tr></table>\r\n\r\n\r\n</div>\r\n<div id=\"comment_10776497\" class=\"cm\">\r\n</div>\r\n\r\n<div id=\"post_rate_div_10776497\"></div>\r\n</div>\r\n</div>\r\n\r\n</td></tr>\r\n<tr><td class=\"plc plm\">\r\n\r\n<div id=\"p_btn\" class=\"mtw mbm hm cl\">\r\n</div>\r\n\r\n</td>\r\n</tr>\r\n<tr id=\"_postposition10776497\"></tr>\r\n<tr>\r\n<td class=\"pls\"></td>\r\n<td class=\"plc\" style=\"overflow:visible;\">\r\n<div class=\"po hin\">\r\n<div class=\"pob cl\">\r\n<em>\r\n<a class=\"fastre\" href=\"forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1291126&amp;repquote=10776497&amp;extra=page%3D1&amp;page=1\" onclick=\"showWindow('reply', this.href)\">回复</a>\r\n</em>\r\n\r\n<p>\r\n<a href=\"javascript:;\" id=\"mgc_post_10776497\" onmouseover=\"showMenu(this.id)\" class=\"showmenu\">使用道具</a>\r\n<a href=\"javascript:;\" onclick=\"showWindow('miscreport10776497', 'misc.php?mod=report&rtype=post&rid=10776497&tid=1291126&fid=36', 'get', -1);return false;\">举报</a>\r\n</p>\r\n\r\n<ul id=\"mgc_post_10776497_menu\" class=\"p_pop mgcmn\" style=\"display: none;\">\r\n</ul>\r\n<script type=\"text/javascript\" reload=\"1\">checkmgcmn('post_10776497')</script>\r\n</div>\r\n</div>\r\n</td>\r\n</tr>\r\n<tr class=\"ad\">\r\n<td class=\"pls\">\r\n</td>\r\n<td class=\"plc\">\r\n</td>\r\n</tr>\r\n</table>\r\n</div>\r\n\r\n<div id=\"post_10778365\" ><table id=\"pid10778365\" class=\"plhin\" summary=\"pid10778365\" cellspacing=\"0\" cellpadding=\"0\">\r\n<tr>\r\n <td class=\"pls\" rowspan=\"2\">\r\n<div id=\"favatar10778365\" class=\"pls cl favatar\">\r\n<div class=\"pi\">\r\n<div class=\"authi\"><a href=\"space-uid-449173.html\" target=\"_blank\" class=\"xw1\">金银先生</a>\r\n</div>\r\n</div>\r\n<div class=\"p_pop blk bui card_gender_0\" id=\"userinfo10778365\" style=\"display: none; margin-top: -11px;\">\r\n<div class=\"m z\">\r\n<div id=\"userinfo10778365_ma\"></div>\r\n</div>\r\n<div class=\"i y\">\r\n<div>\r\n<strong><a href=\"space-uid-449173.html\" target=\"_blank\" class=\"xi2\">金银先生</a></strong>\r\n<em>当前在线</em>\r\n</div><dl class=\"cl\">\r\n<dt>积分</dt><dd><a href=\"home.php?mod=space&uid=449173&do=profile\" target=\"_blank\" class=\"xi2\">104</a></dd>\r\n</dl><div class=\"imicn\">\r\n<a href=\"home.php?mod=space&amp;uid=449173&amp;do=profile\" target=\"_blank\" title=\"查看详细资料\"><img src=\"static/image/common/userinfo.gif\" alt=\"查看详细资料\" /></a>\r\n</div>\r\n<div id=\"avatarfeed\"><span id=\"threadsortswait\"></span></div>\r\n</div>\r\n</div>\r\n<div>\r\n<div class=\"avatar\" onmouseover=\"showauthor(this, 'userinfo10778365')\"><a href=\"space-uid-449173.html\" class=\"avtm\" target=\"_blank\"><img src=\"/uc_server/data/avatar/000/44/91/73_avatar_middle.jpg\" onerror=\"this.onerror=null;this.src='/uc_server/images/noavatar_middle.gif'\" /></a></div>\r\n</div>\r\n<div class=\"tns xg2\"><table cellspacing=\"0\" cellpadding=\"0\"><th><p><a href=\"home.php?mod=space&uid=449173&do=profile\" class=\"xi2\">3</a></p>金钱</th><th><p><a href=\"home.php?mod=space&uid=449173&do=profile\" class=\"xi2\">0</a></p>色币</th><td><p><a href=\"home.php?mod=space&uid=449173&do=profile\" class=\"xi2\">0</a></p>评分</td></table></div>\r\n\r\n<p><em><a href=\"home.php?mod=spacecp&amp;ac=usergroup&amp;gid=14\" target=\"_blank\">Lv5 小有名气</a></em></p>\r\n\r\n\r\n<p><span class=\"pbg2\"  id=\"upgradeprogress_10778365\" onmouseover=\"showMenu({'ctrlid':this.id, 'pos':'12!', 'menuid':'g_up10778365_menu'});\"><span class=\"pbr2\" style=\"width:4%;\"></span></span></p>\r\n<div id=\"g_up10778365_menu\" class=\"tip tip_4\" style=\"display: none;\"><div class=\"tip_horn\"></div><div class=\"tip_c\">Lv5 小有名气, 积分 104, 距离下一级还需 96 积分</div></div>\r\n\r\n<dl class=\"pil cl\">\r\n  <dt>积分</dt><dd><a href=\"home.php?mod=space&uid=449173&do=profile\" target=\"_blank\" class=\"xi2\">104</a></dd>\r\n</dl>\r\n\r\n\r\n<ul class=\"xl xl2 o cl\">\r\n<li class=\"pm2\"><a href=\"home.php?mod=spacecp&amp;ac=pm&amp;op=showmsg&amp;handlekey=showmsg_449173&amp;touid=449173&amp;pmid=0&amp;daterange=2&amp;pid=10778365&amp;tid=1291126\" onclick=\"showWindow('sendpm', this.href);\" title=\"发消息\" class=\"xi2\">发消息</a></li>\r\n</ul>\r\n</div>\r\n</td>\r\n<td class=\"plc\">\r\n<div class=\"pi\">\r\n<strong>\r\n<a href=\"forum.php?mod=redirect&goto=findpost&ptid=1291126&pid=10778365\"   id=\"postnum10778365\" onclick=\"setCopy(this.href, '帖子地址复制成功');return false;\">\r\n板凳</a>\r\n</strong>\r\n<div class=\"pti\">\r\n<div class=\"pdbt\">\r\n</div>\r\n<div class=\"authi\">\r\n<img class=\"authicn vm\" id=\"authicon10778365\" src=\"static/image/common/online_member.gif\" />\r\n<em id=\"authorposton10778365\">发表于 <span title=\"2023-04-28 12:43:06\">昨天&nbsp;12:43</span></em>\r\n<span class=\"xg1\">来自手机</span>\r\n<span class=\"pipe\">|</span>\r\n<a href=\"forum.php?mod=viewthread&amp;tid=1291126&amp;page=1&amp;authorid=449173\" rel=\"nofollow\">只看该作者</a>\r\n</div>\r\n</div>\r\n</div><div class=\"pct\"><div class=\"pcb\">\r\n<div class=\"t_fsz\">\r\n<table cellspacing=\"0\" cellpadding=\"0\"><tr><td class=\"t_f\" id=\"postmessage_10778365\">\r\n身材不错</td></tr></table>\r\n\r\n\r\n</div>\r\n<div id=\"comment_10778365\" class=\"cm\">\r\n</div>\r\n\r\n<div id=\"post_rate_div_10778365\"></div>\r\n</div>\r\n</div>\r\n\r\n</td></tr>\r\n<tr><td class=\"plc plm\">\r\n\r\n<div id=\"p_btn\" class=\"mtw mbm hm cl\">\r\n</div>\r\n\r\n</td>\r\n</tr>\r\n<tr id=\"_postposition10778365\"></tr>\r\n<tr>\r\n<td class=\"pls\"></td>\r\n<td class=\"plc\" style=\"overflow:visible;\">\r\n<div class=\"po hin\">\r\n<div class=\"pob cl\">\r\n<em>\r\n<a class=\"fastre\" href=\"forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1291126&amp;repquote=10778365&amp;extra=page%3D1&amp;page=1\" onclick=\"showWindow('reply', this.href)\">回复</a>\r\n</em>\r\n\r\n<p>\r\n<a href=\"javascript:;\" id=\"mgc_post_10778365\" onmouseover=\"showMenu(this.id)\" class=\"showmenu\">使用道具</a>\r\n<a href=\"javascript:;\" onclick=\"showWindow('miscreport10778365', 'misc.php?mod=report&rtype=post&rid=10778365&tid=1291126&fid=36', 'get', -1);return false;\">举报</a>\r\n</p>\r\n\r\n<ul id=\"mgc_post_10778365_menu\" class=\"p_pop mgcmn\" style=\"display: none;\">\r\n</ul>\r\n<script type=\"text/javascript\" reload=\"1\">checkmgcmn('post_10778365')</script>\r\n</div>\r\n</div>\r\n</td>\r\n</tr>\r\n<tr class=\"ad\">\r\n<td class=\"pls\">\r\n</td>\r\n<td class=\"plc\">\r\n</td>\r\n</tr>\r\n</table>\r\n</div>\r\n\r\n<div class=\"show-text4 pad-tb-10\">\r\n<div class=\"show-info\">\r\n<div class=\"avatar\">AD</div>\r\n</div>\r\n<div class=\"show-main\"><div><a href=\"https://www.d23dx.com/\" class=\"item js-appJump\">\r\n    <h4>鲍鱼盒子，全网高清直播资源‖自拍视频在线看</h4>\r\n    <p>注册免费观看，邀请好友送现金红包。</p>\r\n</a>\r\n<a href=\"https://www.d23dx.com/\" class=\"item js-appJump\">\r\n    <h4>鲍鱼盒子，全网高清直播资源‖自拍视频在线看</h4>\r\n    <p>注册免费观看，邀请好友送现金红包。</p>\r\n</a></div></div>\r\n</div>\r\n<table class=\"plhin\">\r\n<tr class=\"ad\">\r\n<td class=\"pls\"></td>\r\n<td class=\"plc\"></td>\r\n</tr>\r\n</table>\r\n<div id=\"post_10784146\" ><table id=\"pid10784146\" class=\"plhin\" summary=\"pid10784146\" cellspacing=\"0\" cellpadding=\"0\">\r\n<tr>\r\n <a name=\"lastpost\"></a><td class=\"pls\" rowspan=\"2\">\r\n<div id=\"favatar10784146\" class=\"pls cl favatar\">\r\n<div class=\"pi\">\r\n<div class=\"authi\"><a href=\"space-uid-447195.html\" target=\"_blank\" class=\"xw1\">yoscso</a>\r\n</div>\r\n</div>\r\n<div class=\"p_pop blk bui card_gender_0\" id=\"userinfo10784146\" style=\"display: none; margin-top: -11px;\">\r\n<div class=\"m z\">\r\n<div id=\"userinfo10784146_ma\"></div>\r\n</div>\r\n<div class=\"i y\">\r\n<div>\r\n<strong><a href=\"space-uid-447195.html\" target=\"_blank\" class=\"xi2\">yoscso</a></strong>\r\n<em>当前在线</em>\r\n</div><dl class=\"cl\">\r\n<dt>积分</dt><dd><a href=\"home.php?mod=space&uid=447195&do=profile\" target=\"_blank\" class=\"xi2\">1374</a></dd>\r\n</dl><div class=\"imicn\">\r\n<a href=\"home.php?mod=space&amp;uid=447195&amp;do=profile\" target=\"_blank\" title=\"查看详细资料\"><img src=\"static/image/common/userinfo.gif\" alt=\"查看详细资料\" /></a>\r\n</div>\r\n<div id=\"avatarfeed\"><span id=\"threadsortswait\"></span></div>\r\n</div>\r\n</div>\r\n<div>\r\n<div class=\"avatar\" onmouseover=\"showauthor(this, 'userinfo10784146')\"><a href=\"space-uid-447195.html\" class=\"avtm\" target=\"_blank\"><img src=\"/uc_server/data/avatar/000/44/71/95_avatar_middle.jpg\" onerror=\"this.onerror=null;this.src='/uc_server/images/noavatar_middle.gif'\" /></a></div>\r\n</div>\r\n<div class=\"tns xg2\"><table cellspacing=\"0\" cellpadding=\"0\"><th><p><a href=\"home.php?mod=space&uid=447195&do=profile\" class=\"xi2\">101</a></p>金钱</th><th><p><a href=\"home.php?mod=space&uid=447195&do=profile\" class=\"xi2\">0</a></p>色币</th><td><p><a href=\"home.php?mod=space&uid=447195&do=profile\" class=\"xi2\">66</a></p>评分</td></table></div>\r\n\r\n<p><em><a href=\"home.php?mod=spacecp&amp;ac=usergroup&amp;gid=28\" target=\"_blank\">Lv8 武林高手</a></em></p>\r\n\r\n\r\n<p><span class=\"pbg2\"  id=\"upgradeprogress_10784146\" onmouseover=\"showMenu({'ctrlid':this.id, 'pos':'12!', 'menuid':'g_up10784146_menu'});\"><span class=\"pbr2\" style=\"width:37%;\"></span></span></p>\r\n<div id=\"g_up10784146_menu\" class=\"tip tip_4\" style=\"display: none;\"><div class=\"tip_horn\"></div><div class=\"tip_c\">Lv8 武林高手, 积分 1374, 距离下一级还需 626 积分</div></div>\r\n\r\n<dl class=\"pil cl\">\r\n  <dt>积分</dt><dd><a href=\"home.php?mod=space&uid=447195&do=profile\" target=\"_blank\" class=\"xi2\">1374</a></dd>\r\n</dl>\r\n\r\n\r\n<ul class=\"xl xl2 o cl\">\r\n<li class=\"pm2\"><a href=\"home.php?mod=spacecp&amp;ac=pm&amp;op=showmsg&amp;handlekey=showmsg_447195&amp;touid=447195&amp;pmid=0&amp;daterange=2&amp;pid=10784146&amp;tid=1291126\" onclick=\"showWindow('sendpm', this.href);\" title=\"发消息\" class=\"xi2\">发消息</a></li>\r\n</ul>\r\n</div>\r\n</td>\r\n<td class=\"plc\">\r\n<div class=\"pi\">\r\n<strong>\r\n<a href=\"forum.php?mod=redirect&goto=findpost&ptid=1291126&pid=10784146\"   id=\"postnum10784146\" onclick=\"setCopy(this.href, '帖子地址复制成功');return false;\">\r\n地板</a>\r\n</strong>\r\n<div class=\"pti\">\r\n<div class=\"pdbt\">\r\n</div>\r\n<div class=\"authi\">\r\n<img class=\"authicn vm\" id=\"authicon10784146\" src=\"static/image/common/online_member.gif\" />\r\n<em id=\"authorposton10784146\">发表于 <span title=\"2023-04-28 20:51:52\">昨天&nbsp;20:51</span></em>\r\n<span class=\"xg1\">来自手机</span>\r\n<span class=\"pipe\">|</span>\r\n<a href=\"forum.php?mod=viewthread&amp;tid=1291126&amp;page=1&amp;authorid=447195\" rel=\"nofollow\">只看该作者</a>\r\n</div>\r\n</div>\r\n</div><div class=\"pct\"><div class=\"pcb\">\r\n<div class=\"t_fsz\">\r\n<table cellspacing=\"0\" cellpadding=\"0\"><tr><td class=\"t_f\" id=\"postmessage_10784146\">\r\n魔鬼天使真是好棒</td></tr></table>\r\n\r\n\r\n</div>\r\n<div id=\"comment_10784146\" class=\"cm\">\r\n</div>\r\n\r\n<div id=\"post_rate_div_10784146\"></div>\r\n</div>\r\n</div>\r\n\r\n</td></tr>\r\n<tr><td class=\"plc plm\">\r\n\r\n<div id=\"p_btn\" class=\"mtw mbm hm cl\">\r\n</div>\r\n\r\n</td>\r\n</tr>\r\n<tr id=\"_postposition10784146\"></tr>\r\n<tr>\r\n<td class=\"pls\"></td>\r\n<td class=\"plc\" style=\"overflow:visible;\">\r\n<div class=\"po hin\">\r\n<div class=\"pob cl\">\r\n<em>\r\n<a class=\"fastre\" href=\"forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1291126&amp;repquote=10784146&amp;extra=page%3D1&amp;page=1\" onclick=\"showWindow('reply', this.href)\">回复</a>\r\n</em>\r\n\r\n<p>\r\n<a href=\"javascript:;\" id=\"mgc_post_10784146\" onmouseover=\"showMenu(this.id)\" class=\"showmenu\">使用道具</a>\r\n<a href=\"javascript:;\" onclick=\"showWindow('miscreport10784146', 'misc.php?mod=report&rtype=post&rid=10784146&tid=1291126&fid=36', 'get', -1);return false;\">举报</a>\r\n</p>\r\n\r\n<ul id=\"mgc_post_10784146_menu\" class=\"p_pop mgcmn\" style=\"display: none;\">\r\n</ul>\r\n<script type=\"text/javascript\" reload=\"1\">checkmgcmn('post_10784146')</script>\r\n</div>\r\n</div>\r\n</td>\r\n</tr>\r\n<tr class=\"ad\">\r\n<td class=\"pls\">\r\n</td>\r\n<td class=\"plc\">\r\n</td>\r\n</tr>\r\n</table>\r\n</div>\r\n\r\n<div id=\"postlistreply\" class=\"pl\"><div id=\"post_new\" class=\"viewthread_table\" style=\"display: none\"></div></div>\r\n</div>\r\n\r\n\r\n<form method=\"post\" autocomplete=\"off\" name=\"modactions\" id=\"modactions\">\r\n<input type=\"hidden\" name=\"formhash\" value=\"786d433d\" />\r\n<input type=\"hidden\" name=\"optgroup\" />\r\n<input type=\"hidden\" name=\"operation\" />\r\n<input type=\"hidden\" name=\"listextra\" value=\"page%3D1\" />\r\n<input type=\"hidden\" name=\"page\" value=\"1\" />\r\n</form>\r\n\r\n\r\n\r\n<div class=\"pgs mtm mbm cl\">\r\n<span class=\"pgb y\" id=\"visitedforumstmp\" onmouseover=\"$('visitedforums').id = 'visitedforumstmp';this.id = 'visitedforums';showMenu({'ctrlid':this.id,'pos':'21'})\"><a href=\"forum-36-1.html\">返回列表</a></span>\r\n<a id=\"newspecialtmp\" onmouseover=\"$('newspecial').id = 'newspecialtmp';this.id = 'newspecial';showMenu({'ctrlid':this.id})\" onclick=\"showWindow('newthread', 'forum.php?mod=post&action=newthread&fid=36')\" href=\"javascript:;\" title=\"发新帖\"><img src=\"static/image/common/pn_post.png\" alt=\"发新帖\" /></a>\r\n</div>\r\n\r\n<!--[diy=diyfastposttop]--><div id=\"diyfastposttop\" class=\"area\"></div><!--[/diy]-->\r\n<script type=\"text/javascript\">\r\nvar postminchars = parseInt('10');\r\nvar postmaxchars = parseInt('500000000');\r\nvar disablepostctrl = parseInt('0');\r\n</script>\r\n\r\n<div id=\"f_pst\" class=\"pl bm bmw\">\r\n<form method=\"post\" autocomplete=\"off\" id=\"fastpostform\" action=\"forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1291126&amp;extra=page%3D1&amp;replysubmit=yes&amp;infloat=yes&amp;handlekey=fastpost\" onSubmit=\"return fastpostvalidate(this)\">\r\n<table cellspacing=\"0\" cellpadding=\"0\">\r\n<tr>\r\n<td class=\"pls\">\r\n</td>\r\n<td class=\"plc\">\r\n\r\n<span id=\"fastpostreturn\"></span>\r\n\r\n\r\n<div class=\"cl\">\r\n<div id=\"fastsmiliesdiv\" class=\"y\"><div id=\"fastsmiliesdiv_data\"><div id=\"fastsmilies\"></div></div></div><div class=\"hasfsl\" id=\"fastposteditor\">\r\n<div class=\"tedt mtn\">\r\n<div class=\"bar\">\r\n<span class=\"y\">\r\n<a href=\"forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1291126\" onclick=\"return switchAdvanceMode(this.href)\">高级模式</a>\r\n</span><script src=\"static/js/seditor.js?Um1\" type=\"text/javascript\"></script>\r\n<div class=\"fpd\">\r\n<a href=\"javascript:;\" title=\"文字加粗\" class=\"fbld\">B</a>\r\n<a href=\"javascript:;\" title=\"设置文字颜色\" class=\"fclr\" id=\"fastpostforecolor\">Color</a>\r\n<a id=\"fastpostimg\" href=\"javascript:;\" title=\"图片\" class=\"fmg\">Image</a>\r\n<a id=\"fastposturl\" href=\"javascript:;\" title=\"添加链接\" class=\"flnk\">Link</a>\r\n<a id=\"fastpostquote\" href=\"javascript:;\" title=\"引用\" class=\"fqt\">Quote</a>\r\n<a id=\"fastpostcode\" href=\"javascript:;\" title=\"代码\" class=\"fcd\">Code</a>\r\n<a href=\"javascript:;\" class=\"fsml\" id=\"fastpostsml\">Smilies</a>\r\n</div></div>\r\n<div class=\"area\">\r\n<div class=\"pt hm\">\r\n您需要登录后才可以回帖 <a href=\"member.php?mod=logging&amp;action=login\" onclick=\"showWindow('login', this.href)\" class=\"xi2\">登录</a> | <a href=\"member.php?mod=register\" class=\"xi2\">立即注册</a>\r\n</div>\r\n</div>\r\n</div>\r\n</div>\r\n</div>\r\n<div id=\"seccheck_fastpost\">\r\n</div>\r\n\r\n\r\n<input type=\"hidden\" name=\"formhash\" value=\"786d433d\" />\r\n<input type=\"hidden\" name=\"usesig\" value=\"\" />\r\n<input type=\"hidden\" name=\"subject\" value=\"  \" />\r\n<p class=\"ptm pnpost\">\r\n<a href=\"home.php?mod=spacecp&amp;ac=credit&amp;op=rule&amp;fid=36\" class=\"y\" target=\"_blank\">本版积分规则</a>\r\n<button type=\"button\" onclick=\"showWindow('login', 'member.php?mod=logging&action=login&guestmessage=yes')\"  onmouseover=\"checkpostrule('seccheck_fastpost', 'ac=reply');this.onmouseover=null\" name=\"replysubmit\" id=\"fastpostsubmit\" class=\"pn pnc vm\" value=\"replysubmit\" tabindex=\"5\"><strong>发表回复</strong></button>\r\n<label for=\"fastpostrefresh\"><input id=\"fastpostrefresh\" type=\"checkbox\" class=\"pc\" />回帖后跳转到最后一页</label>\r\n<script type=\"text/javascript\">if(getcookie('fastpostrefresh') == 1) {$('fastpostrefresh').checked=true;}</script>\r\n</p>\r\n</td>\r\n</tr>\r\n</table>\r\n</form>\r\n</div>\r\n\r\n<div id=\"visitedforums_menu\" class=\"p_pop blk cl\" style=\"display: none;\">\r\n<table cellspacing=\"0\" cellpadding=\"0\">\r\n<tr>\r\n<td id=\"v_forums\">\r\n<h3 class=\"mbn pbn bbda xg1\">浏览过的版块</h3>\r\n<ul class=\"xl xl1\">\r\n<li><a href=\"forum-103-1.html\">高清中文字幕</a></li><li><a href=\"forum-142-1.html\">转帖交流区</a></li><li><a href=\"forum-159-1.html\">新作区</a></li><li><a href=\"forum-39-1.html\">动漫原创</a></li><li><a href=\"forum-141-1.html\">网友原创区</a></li><li><a href=\"forum-117-1.html\">卡通动漫</a></li><li><a href=\"forum-46-1.html\">剧情三级</a></li><li><a href=\"forum-109-1.html\">中文字幕</a></li><li><a href=\"forum-42-1.html\">日韩无码</a></li></ul>\r\n</td>\r\n</tr>\r\n</table>\r\n</div>\r\n<script type=\"text/javascript\">\r\nnew lazyload();\r\n</script>\r\n<script type=\"text/javascript\">document.onkeyup = function(e){keyPageScroll(e, 0, 0, 'forum.php?mod=viewthread&tid=1291126', 1);}</script>\r\n</div>\r\n\r\n<div class=\"wp mtn\">\r\n<!--[diy=diy3]--><div id=\"diy3\" class=\"area\"></div><!--[/diy]-->\r\n</div>\r\n\r\n<script type=\"text/javascript\">\r\nfunction succeedhandle_followmod(url, msg, values) {\r\nvar fObj = $('followmod_'+values['fuid']);\r\nif(values['type'] == 'add') {\r\nfObj.innerHTML = '不收听';\r\nfObj.href = 'home.php?mod=spacecp&ac=follow&op=del&fuid='+values['fuid'];\r\n} else if(values['type'] == 'del') {\r\nfObj.innerHTML = '收听TA';\r\nfObj.href = 'home.php?mod=spacecp&ac=follow&op=add&hash=068bf686&fuid='+values['fuid'];\r\n}\r\n}\r\nfixed_avatar([10775860,10776497,10778365,10784146], 1);\r\n</script><div class=\"show-text cl\"><a href=\"https://tz112.com\" target=\"_blank\" class=\"js-randomBg\">亚博美女赌场</a>\r\n<a href=\"https://5656456565.com/vote_topic_80615.do\" target=\"_blank\" class=\"js-randomBg\">澳门现金赌场</a>\r\n<a href=\"http://23.225.52.51:4466/vip103.html\" target=\"_blank\" class=\"js-randomBg\">澳门新葡京</a>\r\n<a href=\"https://by301.hdzshbk.cn/?channelCode=by_147\" target=\"_blank\" class=\"js-randomBg\">鲍鱼全球黄播</a>\r\n<a href=\"http://23.224.188.2:2939/vip107.html\" target=\"_blank\" class=\"js-randomBg\">澳门威尼斯人</a>\r\n<a href=\"https://154.82.93.23:443/vip058\" target=\"_blank\" class=\"js-randomBg\">澳门葡京赌场</a>\r\n<a href=\"https://66861056.app/registered\" target=\"_blank\" class=\"js-randomBg\">天天送红包</a>\r\n<a href=\"https://5581268.cc:8443/?shareName=5581268.cc\" target=\"_blank\" class=\"js-randomBg\">【开元棋牌】</a>\r\n<a href=\"https://h1116.cc:15555\" target=\"_blank\" class=\"js-randomBg\">现金棋牌</a>\r\n<a href=\"https://www.rckif.com/001.html\" target=\"_blank\" class=\"js-randomBg\">91成人抖音</a>\r\n</div>\r\n<div class=\"h10\"></div>\r\n\t</div>\r\n<div id=\"ft\" class=\"wp cl\">\r\n<div id=\"flk\" class=\"y\">\r\n<p>\r\n广告联系：sehuatang@gmail.com<span class=\"pipe\">|</span>\r\n<a href=\"forum.php?mod=misc&action=showdarkroom\" >小黑屋</a><span class=\"pipe\">|</span><strong><a href=\"https://www.sehuatang.net\" target=\"_blank\">98堂[原色花堂]</a></strong>\r\n</p>\r\n<p class=\"xs0\">\r\nGMT+8, 2023-04-29 00:24<span id=\"debuginfo\">\r\n, Processed in 0.129361 second(s), Total 5, Slave 4 queries\r\n, Redis On.\r\n</span>\r\n</p>\r\n</div>\r\n<div id=\"frt\">\r\n<p>Powered by <strong><a href=\"http://www.discuz.net\" target=\"_blank\">Discuz!</a></strong> <em>X3.4</em></p>\r\n<p class=\"xs0\">Copyright &copy; 2001-2021, Tencent Cloud.</p>\r\n</div></div>\r\n\r\n<div id=\"scrolltop\">\r\n<span><a href=\"forum.php?mod=post&amp;action=reply&amp;fid=36&amp;tid=1291126&amp;extra=page%3D1&amp;page=1\" onclick=\"showWindow('reply', this.href)\" class=\"replyfast\" title=\"快速回复\"><b>快速回复</b></a></span>\r\n<span hidefocus=\"true\"><a title=\"返回顶部\" onclick=\"window.scrollTo('0','0')\" class=\"scrolltopa\" ><b>返回顶部</b></a></span>\r\n<span>\r\n<a href=\"forum-36-1.html\" hidefocus=\"true\" class=\"returnlist\" title=\"返回列表\"><b>返回列表</b></a>\r\n</span>\r\n</div>\r\n<script type=\"text/javascript\">_attachEvent(window, 'scroll', function () { showTopLink(); });checkBlind();randomBg('js-randomBg');</script>\r\n<script type=\"text/javascript\">$(\"debuginfo\") ? $(\"debuginfo\").innerHTML = \", Updated at 2023-04-29 00:24:23, Processed in 0.003539 second(s).\" : \"\";</script></body></html>"
	config := "{\"name_content\":\"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/table[1]/tbody/tr/td[@class='plc ptm pbn vwthd']/h1[@class='ts']/span[@id='thread_subject']\",\"taq_content\":\"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/table[1]/tbody/tr/td[@class='plc ptm pbn vwthd']/h1[@class='ts']/a\",\"pubishTime_content\":\"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/div[@id]/table[@id]/tbody/tr[1]/td[@class='plc']/div[@class='pi']/div[@class='pti']/div[@class='authi']/em[@id]\",\"start_num_content\":\"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/div[@id]/table[@id]/tbody/tr[2]/td[@class='plc plm']/div[@id='p_btn']/a[@id='k_favorite']/i/span[@id='favoritenumber']\",\"discussNum_content\":\"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/table[1]/tbody/tr/td[@class='pls ptn pbn']/div[@class='hm ptn']/span[@class='xi1'][2]\",\"pv_content\":\"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/table[1]/tbody/tr/td[@class='pls ptn pbn']/div[@class='hm ptn']/span[@class='xi1'][1]\",\"evaluate_content\":\"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/div[@id]/table[@id]/tbody/tr[1]/td[@class='plc']/div[@class='pct']/div[@class='pcb']/dl[@id]/dd/table[@class='ratl']/tbody[1]/tr/th[@class='xw1']\",\"pubisher\":{\"name_content\":\"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/div[@id]/table[@id]/tbody/tr[1]/td[@class='pls']/div[@id]/div[@class='pi']/div[@class='authi']/a[@class='xw1']\",\"money_content\":\"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/div[@id]/table[@id]/tbody/tr[1]/td[@class='pls']/div[@id]/div[@class='tns xg2']/table/tbody/tr/th[1]/p/a[@class='xi2']\",\"sebi_content\":\"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/div[@id]/table[@id]/tbody/tr[1]/td[@class='pls']/div[@id]/div[@class='tns xg2']/table/tbody/tr/th[2]/p/a[@class='xi2']\",\"evaluateScore_content\":\"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/div[@id]/table[@id]/tbody/tr[1]/td[@class='pls']/div[@id]/div[@class='tns xg2']/table/tbody/tr/td/p/a[@class='xi2']\",\"score_content\":\"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/div[@id]/table[@id]/tbody/tr[1]/td[@class='pls']/div[@id]/dl[@class='pil cl']/dd/a[@class='xi2']\"},\"content_content\":\"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/div[@id]/table[@id]/tbody/tr[1]/td[@class='plc']/div[@class='pct']/div[@class='pcb']/div[@class='t_fsz']/table/tbody/tr/td[@id]\",\"magnet_content\":\"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/div[@id]/table[@id]/tbody/tr[1]/td[@class='plc']/div[@class='pct']/div[@class='pcb']/div[@class='t_fsz']/table/tbody/tr/td[@id]/div[@class='blockcode']/div[@id]/ol/li\",\"file_attr\":[\"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/div[@id]/table[@id]/tbody/tr[1]/td[@class='plc']/div[@class='pct']/div[@class='pcb']/div[@class='t_fsz']/table/tbody/tr/td[@id]/ignore_js_op/img[@id]\"]}"
	srt := ""
	fmt.Println(Exe(config, srt, html, ""))
}

func TestHtml01(t *testing.T) {
	html := "<!DOCTYPE html>\n<html>\n<head>\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" />\n<title>ADN-315 女好きの親父は僕の妻をいやらしい目で見ていたんです...。 竹内夏希 - 转帖交流区 -  98堂[原色花堂] -  Powered by Discuz!</title>\n<link href=\"/thread-1293441-1-1.html\" rel=\"canonical\" />\n<meta name=\"keywords\" content=\"ADN-315 女好きの親父は僕の妻をいやらしい目で見ていたんです...。 竹内夏希\" />\n<meta name=\"description\" content=\"【资源名称】：ADN-315 女好きの親父は僕の妻をいやらしい目で見ていたんです...。 竹内夏希【文件大小】：971MB【下载地址】：B263FBA45B7592B527FB60A8551D09415BB7B531【相关密码】：无【图片预览】： ... ADN-315 女好きの親父は僕の妻をいやらしい目で見ていたんです...。 竹内夏希 ,98堂[原色花堂]\" />\n<meta name=\"generator\" content=\"Discuz! X3.4\" />\n<meta name=\"author\" content=\"Discuz! Team and Comsenz UI Team\" />\n<meta name=\"copyright\" content=\"2001-2021 Tencent Cloud.\" />\n<meta name=\"MSSmartTagsPreventParsing\" content=\"True\" />\n<meta http-equiv=\"MSThemeCompatible\" content=\"Yes\" />\n<base href=\"/\" /><link rel=\"stylesheet\" type=\"text/css\" href=\"data/cache/style_1_common.css?Um1\" /><link rel=\"stylesheet\" type=\"text/css\" href=\"data/cache/style_1_forum_viewthread.css?Um1\" /><link rel=\"stylesheet\" id=\"css_extstyle\" type=\"text/css\" href=\"./template/default/style/t1/style.css\" /><script type=\"text/javascript\">var STYLEID = '1', STATICURL = 'static/', IMGDIR = 'static/image/common', VERHASH = 'Um1', charset = 'utf-8', discuz_uid = '0', cookiepre = 'cPNj_2132_', cookiedomain = '', cookiepath = '/', showusercard = '1', attackevasive = '0', disallowfloat = 'newthread', creditnotice = '3|积分|,5|金钱|,7|色币|,8|评分|', defaultstyle = './template/default/style/t1', REPORTURL = 'aHR0cHM6Ly93d3cuMThzdG0uY24vdGhyZWFkLTEyOTM0NDEtMS0xLmh0bWw=', SITEURL = '/', JSPATH = 'static/js/', CSSPATH = 'data/cache/style_', DYNAMICURL = '';</script>\n<script src=\"static/js/common.js?Um1\" type=\"text/javascript\"></script>\n<script src=\"static/libs/fingerprintjs/fp.min.js\" type=\"text/javascript\" onload=\"fingerprint()\"></script>\n<meta name=\"application-name\" content=\"98堂[原色花堂]\" />\n<meta name=\"msapplication-tooltip\" content=\"98堂[原色花堂]\" />\n<meta name=\"msapplication-task\" content=\"name=新帖瀑布流;action-uri=/portal.php;icon-uri=static/image/common/portal.ico\" /><meta name=\"msapplication-task\" content=\"name=论坛;action-uri=/forum.php;icon-uri=static/image/common/bbs.ico\" />\n<link rel=\"archives\" title=\"98堂[原色花堂]\" href=\"/archiver/\" />\n<script src=\"static/js/forum.js?Um1\" type=\"text/javascript\"></script>\n</head>\n\n<body id=\"nv_forum\" class=\"pg_viewthread\" onkeydown=\"if(event.keyCode==27) return false;\">\n<div id=\"append_parent\"></div><div id=\"ajaxwaitid\"></div>\n<div id=\"toptb\" class=\"cl\">\n<div class=\"wp\">\n<div class=\"z\"><a href=\"javascript:;\"  onclick=\"setHomepage('/');\">设为首页</a><a href=\"/\"  onclick=\"addFavorite(this.href, '98堂[原色花堂]');return false;\">收藏本站</a><script src=\"static/js/language.js\" type=\"text/javascript\"></script>\n</div>\n<div class=\"y\">\n<a id=\"switchblind\" href=\"javascript:;\" onclick=\"toggleBlind(this)\" title=\"开启辅助访问\" class=\"switchblind\"></a>\n<a href=\"javascript:;\" id=\"switchwidth\" onclick=\"widthauto(this)\" title=\"切换到宽版\" class=\"switchwidth\">切换到宽版</a>\n</div>\n</div>\n</div>\n\n<div id=\"qmenu_menu\" class=\"p_pop blk\" style=\"display: none;\">\n<div class=\"ptm pbw hm\">\n请 <a href=\"javascript:;\" class=\"xi2\" onclick=\"lsSubmit()\"><strong>登录</strong></a> 后使用快捷导航<br />没有帐号？<a href=\"member.php?mod=register\" class=\"xi2 xw1\">立即注册</a>\n</div>\n<div id=\"fjump_menu\" class=\"btda\"></div></div><div id=\"hd\">\n<div class=\"wp\">\n<div class=\"hdc cl\"><h2><a href=\"./\" title=\"98堂[原色花堂]\"><img src=\"static/image/common/logo.png\" alt=\"98堂[原色花堂]\" border=\"0\" /></a></h2><script src=\"static/js/logging.js?Um1\" type=\"text/javascript\"></script>\n<form method=\"post\" autocomplete=\"off\" id=\"lsform\" action=\"/member.php?mod=logging&amp;action=login&amp;loginsubmit=yes&amp;infloat=yes&amp;lssubmit=yes\" onsubmit=\"return lsSubmit();\">\n<div class=\"fastlg cl\">\n<span id=\"return_ls\" style=\"display:none\"></span>\n<div class=\"y pns\">\n<table cellspacing=\"0\" cellpadding=\"0\">\n<tr>\n<td>\n<span class=\"ftid\">\n<select name=\"fastloginfield\" id=\"ls_fastloginfield\" width=\"40\" tabindex=\"900\">\n<option value=\"username\">用户名</option>\n<option value=\"email\">Email</option>\n</select>\n</span>\n<script type=\"text/javascript\">simulateSelect('ls_fastloginfield')</script>\n</td>\n<td><input type=\"text\" name=\"username\" id=\"ls_username\" autocomplete=\"off\" class=\"px vm\" tabindex=\"901\" /></td>\n<td class=\"fastlg_l\"><label for=\"ls_cookietime\"><input type=\"checkbox\" name=\"cookietime\" id=\"ls_cookietime\" class=\"pc\" value=\"2592000\" tabindex=\"903\" checked=\"checked\"/>自动登录</label></td>\n<td>&nbsp;<a href=\"javascript:;\" onclick=\"showWindow('login', 'member.php?mod=logging&action=login&viewlostpw=1')\">找回密码</a></td>\n</tr>\n<tr>\n<td><label for=\"ls_password\" class=\"z psw_w\">密码</label></td>\n<td><input type=\"password\" name=\"password\" id=\"ls_password\" class=\"px vm\" autocomplete=\"off\" tabindex=\"902\" /></td>\n<td class=\"fastlg_l\"><button type=\"submit\" class=\"pn vm\" tabindex=\"904\" style=\"width: 75px;\"><em>登录</em></button></td>\n<td>&nbsp;<a href=\"member.php?mod=register\" class=\"xi2 xw1\">立即注册</a></td>\n</tr>\n</table>\n<input type=\"hidden\" name=\"formhash\" value=\"786d433d\" />\n<input type=\"hidden\" name=\"quickforward\" value=\"yes\" />\n<input type=\"hidden\" name=\"handlekey\" value=\"ls\" />\n</div>\n</div>\n</form>\n\n</div>\n\n<div id=\"nv\">\n<a href=\"javascript:;\" id=\"qmenu\" onmouseover=\"delayShow(this, function () {showMenu({'ctrlid':'qmenu','pos':'34!','ctrlclass':'a','duration':2});showForummenu(142);})\">快捷导航</a>\n<ul><li class=\"a\" id=\"mn_forum\" ><a href=\"forum.php\" hidefocus=\"true\" title=\"BBS\"  >论坛<span>BBS</span></a></li><li id=\"mn_N743b\" ><a href=\"book.html\" hidefocus=\"true\"  >小说</a></li><li id=\"mn_portal\" ><a href=\"portal.php\" hidefocus=\"true\" title=\"Portal\"  >新帖瀑布流<span>Portal</span></a></li></ul>\n</div>\n<div class=\"p_pop h_pop\" id=\"mn_userapp_menu\" style=\"display: none\"></div><div id=\"mu\" class=\"cl\">\n</div><div id=\"scbar\" class=\"cl\">\n<form id=\"scbar_form\" method=\"post\" autocomplete=\"off\" onsubmit=\"searchFocus($('scbar_txt'))\" action=\"search.php?searchsubmit=yes\" target=\"_blank\">\n<input type=\"hidden\" name=\"mod\" id=\"scbar_mod\" value=\"search\" />\n<input type=\"hidden\" name=\"formhash\" value=\"786d433d\" />\n<input type=\"hidden\" name=\"srchtype\" value=\"title\" />\n<input type=\"hidden\" name=\"srhfid\" value=\"142\" />\n<input type=\"hidden\" name=\"srhlocality\" value=\"forum::viewthread\" />\n<table cellspacing=\"0\" cellpadding=\"0\">\n<tr>\n<td class=\"scbar_icon_td\"></td>\n<td class=\"scbar_txt_td\"><input type=\"text\" name=\"srchtxt\" id=\"scbar_txt\" value=\"请输入搜索内容\" autocomplete=\"off\" x-webkit-speech speech /></td>\n<td class=\"scbar_type_td\"><a href=\"javascript:;\" id=\"scbar_type\" class=\"xg1\" onclick=\"showMenu(this.id)\" hidefocus=\"true\">搜索</a></td>\n<td class=\"scbar_btn_td\"><button type=\"submit\" name=\"searchsubmit\" id=\"scbar_btn\" sc=\"1\" class=\"pn pnc\" value=\"true\"><strong class=\"xi2\">搜索</strong></button></td>\n<td class=\"scbar_hot_td\">\n<div id=\"scbar_hot\">\n</div>\n</td>\n</tr>\n</table>\n</form>\n</div>\n<ul id=\"scbar_type_menu\" class=\"p_pop\" style=\"display: none;\"><li><a href=\"javascript:;\" rel=\"curforum\" fid=\"142\" >本版</a></li><li><a href=\"javascript:;\" rel=\"forum\" class=\"curtype\">帖子</a></li><li><a href=\"javascript:;\" rel=\"user\">用户</a></li></ul>\n<script type=\"text/javascript\">\ninitSearchmenu('scbar', '');\n</script>\n</div>\n</div>\n\n\n<div id=\"wp\" class=\"wp\">\n<div class=\"h10\"></div>\n<div class=\"show-text cl\">\n        <a href=\"https://tz112.com\" target=\"_blank\" class=\"js-randomBg\">亚博美女赌场</a>\n        <a href=\"https://5656456565.com/vote_topic_80615.do\" target=\"_blank\" class=\"js-randomBg\">澳门现金赌场</a>\n        <a href=\"http://23.225.52.51:4466/vip103.html\" target=\"_blank\" class=\"js-randomBg\">澳门新葡京</a>\n        <a href=\"https://by301.hdzshbk.cn/?channelCode=by_147\" target=\"_blank\" class=\"js-randomBg\">鲍鱼全球黄播</a>\n        <a href=\"http://23.224.188.2:2939/vip107.html\" target=\"_blank\" class=\"js-randomBg\">澳门威尼斯人</a>\n        <a href=\"https://154.82.93.23:443/vip058\" target=\"_blank\" class=\"js-randomBg\">澳门葡京赌场</a>\n        <a href=\"https://66861056.app/registered\" target=\"_blank\" class=\"js-randomBg\">天天送红包</a>\n        <a href=\"https://5581268.cc:8443/?shareName=5581268.cc\" target=\"_blank\" class=\"js-randomBg\">【开元棋牌】</a>\n        <a href=\"https://h1116.cc:15555\" target=\"_blank\" class=\"js-randomBg\"><font color=\"#00ff00\">现金棋牌</font></a>\n        <a href=\"https://www.rckif.com/001.html\" target=\"_blank\" class=\"js-randomBg\">91成人抖音</a>\n    </div>\n<script type=\"text/javascript\">var fid = parseInt('142'), tid = parseInt('1293441');</script>\n\n<script src=\"static/js/forum_viewthread.js?Um1\" type=\"text/javascript\"></script>\n<script type=\"text/javascript\">zoomstatus = parseInt(1);var imagemaxwidth = '600';var aimgcount = new Array();</script>\n\n<style id=\"diy_style\" type=\"text/css\"></style>\n<!--[diy=diynavtop]--><div id=\"diynavtop\" class=\"area\"></div><!--[/diy]-->\n<div id=\"pt\" class=\"bm cl\">\n<div class=\"z\">\n<a href=\"./\" class=\"nvhm\" title=\"首页\">98堂[原色花堂]</a><em>&raquo;</em><a href=\"forum.php\">论坛</a> <em>&rsaquo;</em> <a href=\"forum.php?gid=94\">综合讨论区</a> <em>&rsaquo;</em> <a href=\"forum-142-1.html\">转帖交流区</a> <em>&rsaquo;</em> <a href=\"thread-1293441-1-1.html\">ADN-315 女好きの親父は僕の妻をいやらしい目で見ていた ...</a>\n</div>\n</div>\n\n<style id=\"diy_style\" type=\"text/css\"></style>\n<div class=\"wp\">\n<!--[diy=diy1]--><div id=\"diy1\" class=\"area\"></div><!--[/diy]-->\n</div>\n\n<div id=\"ct\" class=\"wp cl\">\n<div id=\"pgt\" class=\"pgs mbm cl \">\n<div class=\"pgt\"></div>\n<span class=\"y pgb\" id=\"visitedforums\" onmouseover=\"$('visitedforums').id = 'visitedforumstmp';this.id = 'visitedforums';showMenu({'ctrlid':this.id,'pos':'34'})\"><a href=\"forum-142-1.html\">返回列表</a></span>\n<a id=\"newspecial\" onmouseover=\"$('newspecial').id = 'newspecialtmp';this.id = 'newspecial';showMenu({'ctrlid':this.id})\" onclick=\"showWindow('newthread', 'forum.php?mod=post&action=newthread&fid=142')\" href=\"javascript:;\" title=\"发新帖\"><img src=\"static/image/common/pn_post.png\" alt=\"发新帖\" /></a></div>\n\n\n\n<div id=\"postlist\" class=\"pl bm\">\n<table cellspacing=\"0\" cellpadding=\"0\">\n<tr>\n<td class=\"pls ptn pbn\">\n<div class=\"hm ptn\">\n<span class=\"xg1\">查看:</span> <span class=\"xi1\">114</span><span class=\"pipe\">|</span><span class=\"xg1\">回复:</span> <span class=\"xi1\">0</span>\n</div>\n</td>\n<td class=\"plc ptm pbn vwthd\">\n<div class=\"y\">\n<a href=\"forum.php?mod=viewthread&amp;action=printable&amp;tid=1293441\" title=\"打印\" target=\"_blank\"><img src=\"static/image/common/print.png\" alt=\"打印\" class=\"vm\" /></a>\n<a href=\"forum.php?mod=redirect&amp;goto=nextoldset&amp;tid=1293441\" title=\"上一主题\"><img src=\"static/image/common/thread-prev.png\" alt=\"上一主题\" class=\"vm\" /></a>\n<a href=\"forum.php?mod=redirect&amp;goto=nextnewset&amp;tid=1293441\" title=\"下一主题\"><img src=\"static/image/common/thread-next.png\" alt=\"下一主题\" class=\"vm\" /></a>\n</div>\n<h1 class=\"ts\">\n<a href=\"forum.php?mod=forumdisplay&amp;fid=142&amp;filter=typeid&amp;typeid=700\">[亚洲有码]</a>\n<span id=\"thread_subject\">ADN-315 女好きの親父は僕の妻をいやらしい目で見ていたんです...。 竹内夏希</span>\n</h1>\n<span class=\"xg1\">\n<a href=\"thread-1293441-1-1.html\" onclick=\"return copyThreadUrl(this, '98堂[原色花堂]')\" >[复制链接]</a>\n</span>\n</td>\n</tr>\n</table>\n\n\n<table cellspacing=\"0\" cellpadding=\"0\" class=\"ad\">\n<tr>\n<td class=\"pls\">\n</td>\n<td class=\"plc\">\n</td>\n</tr>\n</table><style>\n.show-text4{position:relative;padding-left:170px;min-height:52px}\n.show-text4 .show-info{position:absolute;top:0;left:0;width:160px;height:100%;text-align:center;display:block;border-right:1px solid #c2d5e3}\n.show-text4 .show-info .avatar{position:absolute;top:50%;left:50%;margin-top:-23px;margin-left:-23px;width:46px;height:46px;line-height:46px;-webkit-border-radius:5rem;border-radius:5rem;background:#c1c1c1;color:#fff;text-align:center;font-size:18px}\n.show-text4 .show-main{padding:5px 10px}\n.show-text4 .item{width:50%;float:left;display:block;text-decoration:none}\n.show-text4 .item h4{font-size:16px;color:#36f}\n.show-text4 .item p{color:#777}\n</style><div id=\"post_10794090\" ><table id=\"pid10794090\" class=\"plhin\" summary=\"pid10794090\" cellspacing=\"0\" cellpadding=\"0\">\n<tr>\n<a name=\"newpost\"></a> <a name=\"lastpost\"></a><td class=\"pls\" rowspan=\"2\">\n<div id=\"favatar10794090\" class=\"pls cl favatar\">\n<div class=\"pi\">\n<div class=\"authi\"><a href=\"space-uid-444819.html\" target=\"_blank\" class=\"xw1\">guihuagao</a>\n</div>\n</div>\n<div class=\"p_pop blk bui card_gender_0\" id=\"userinfo10794090\" style=\"display: none; margin-top: -11px;\">\n<div class=\"m z\">\n<div id=\"userinfo10794090_ma\"></div>\n</div>\n<div class=\"i y\">\n<div>\n<strong><a href=\"space-uid-444819.html\" target=\"_blank\" class=\"xi2\">guihuagao</a></strong>\n<em>当前在线</em>\n</div><dl class=\"cl\">\n<dt>积分</dt><dd><a href=\"home.php?mod=space&uid=444819&do=profile\" target=\"_blank\" class=\"xi2\">553</a></dd>\n</dl><div class=\"imicn\">\n<a href=\"home.php?mod=space&amp;uid=444819&amp;do=profile\" target=\"_blank\" title=\"查看详细资料\"><img src=\"static/image/common/userinfo.gif\" alt=\"查看详细资料\" /></a>\n</div>\n<div id=\"avatarfeed\"><span id=\"threadsortswait\"></span></div>\n</div>\n</div>\n<div>\n<div class=\"avatar\" onmouseover=\"showauthor(this, 'userinfo10794090')\"><a href=\"space-uid-444819.html\" class=\"avtm\" target=\"_blank\"><img src=\"/uc_server/data/avatar/000/44/48/19_avatar_middle.jpg\" onerror=\"this.onerror=null;this.src='/uc_server/images/noavatar_middle.gif'\" /></a></div>\n</div>\n<div class=\"tns xg2\"><table cellspacing=\"0\" cellpadding=\"0\"><th><p><a href=\"home.php?mod=space&uid=444819&do=profile\" class=\"xi2\">231</a></p>金钱</th><th><p><a href=\"home.php?mod=space&uid=444819&do=profile\" class=\"xi2\">0</a></p>色币</th><td><p><a href=\"home.php?mod=space&uid=444819&do=profile\" class=\"xi2\">12</a></p>评分</td></table></div>\n\n<p><em><a href=\"home.php?mod=spacecp&amp;ac=usergroup&amp;gid=27\" target=\"_blank\">Lv7 四方游侠</a></em></p>\n\n\n<p><span class=\"pbg2\"  id=\"upgradeprogress_10794090\" onmouseover=\"showMenu({'ctrlid':this.id, 'pos':'12!', 'menuid':'g_up10794090_menu'});\"><span class=\"pbr2\" style=\"width:10%;\"></span></span></p>\n<div id=\"g_up10794090_menu\" class=\"tip tip_4\" style=\"display: none;\"><div class=\"tip_horn\"></div><div class=\"tip_c\">Lv7 四方游侠, 积分 553, 距离下一级还需 447 积分</div></div>\n\n<dl class=\"pil cl\">\n  <dt>积分</dt><dd><a href=\"home.php?mod=space&uid=444819&do=profile\" target=\"_blank\" class=\"xi2\">553</a></dd>\n</dl>\n\n\n<ul class=\"xl xl2 o cl\">\n<li class=\"pm2\"><a href=\"home.php?mod=spacecp&amp;ac=pm&amp;op=showmsg&amp;handlekey=showmsg_444819&amp;touid=444819&amp;pmid=0&amp;daterange=2&amp;pid=10794090&amp;tid=1293441\" onclick=\"showWindow('sendpm', this.href);\" title=\"发消息\" class=\"xi2\">发消息</a></li>\n</ul>\n</div>\n</td>\n<td class=\"plc\">\n<div class=\"pi\">\n<div id=\"fj\" class=\"y\">\n<label class=\"z\">电梯直达</label>\n<input type=\"text\" class=\"px p_fre z\" size=\"2\" onkeyup=\"$('fj_btn').href='forum.php?mod=redirect&ptid=1293441&authorid=0&postno='+this.value\" onkeydown=\"if(event.keyCode==13) {window.location=$('fj_btn').href;return false;}\" title=\"跳转到指定楼层\" />\n<a href=\"javascript:;\" id=\"fj_btn\" class=\"z\" title=\"跳转到指定楼层\"><img src=\"static/image/common/fj_btn.png\" alt=\"跳转到指定楼层\" class=\"vm\" /></a>\n</div>\n<strong>\n<a href=\"thread-1293441-1-1.html\"   id=\"postnum10794090\" onclick=\"setCopy(this.href, '帖子地址复制成功');return false;\">\n楼主</a>\n</strong>\n<div class=\"pti\">\n<div class=\"pdbt\">\n</div>\n<div class=\"authi\">\n<img class=\"authicn vm\" id=\"authicon10794090\" src=\"static/image/common/online_member.gif\" />\n<em id=\"authorposton10794090\">发表于 <span title=\"2023-04-29 16:25:28\">1&nbsp;小时前</span></em>\n<span class=\"pipe\">|</span>\n<a href=\"forum.php?mod=viewthread&amp;tid=1293441&amp;page=1&amp;authorid=444819\" rel=\"nofollow\">只看该作者</a>\n<span class=\"pipe\">|</span><a href=\"forum.php?mod=viewthread&amp;tid=1293441&amp;from=album\">只看大图</a>\n<span class=\"none\"><img src=\"static/image/common/arw_r.gif\" class=\"vm\" alt=\"回帖奖励\" /></span>\n<span class=\"pipe show\">|</span><a href=\"forum.php?mod=viewthread&amp;tid=1293441&amp;extra=page%3D1&amp;ordertype=1\"  class=\"show\">倒序浏览</a>\n<span class=\"pipe show\">|</span><a href=\"javascript:;\" onclick=\"readmode($('thread_subject').innerHTML, 10794090);\" class=\"show\">阅读模式</a>\n</div>\n</div>\n</div><div class=\"pct\"><style type=\"text/css\">.pcb{margin-right:0}</style><div class=\"pcb\">\n<div class=\"t_fsz\">\n<table cellspacing=\"0\" cellpadding=\"0\"><tr><td class=\"t_f\" id=\"postmessage_10794090\">\n【资源名称】：ADN-315 女好きの親父は僕の妻をいやらしい目で見ていたんです...。 竹内夏希<br />\n【文件大小】：971MB<br />\n【下载地址】：B263FBA45B7592B527FB60A8551D09415BB7B531<br />\n【相关密码】：无<br />\n【图片预览】：<br />\n<br />\n</td></tr></table>\n\n<div class=\"pattl\">\n\n\n<ignore_js_op>\n\n<dl class=\"tattl attm\">\n<dt></dt>\n<dd>\n\n<p class=\"mbn\">\n<a href=\"forum.php?mod=attachment&amp;aid=NDcwNDcyNHwzNTU1NTkwYnwxNjgyNzYxNDAwfDB8MTI5MzQ0MQ%3D%3D&amp;nothumb=yes\" onmouseover=\"showMenu({'ctrlid':this.id,'pos':'12'})\" id=\"aid4704724\" class=\"xw1\" target=\"_blank\">mmexport1682756678137.jpg</a>\n<em class=\"xg1\">(173.67 KB, 下载次数: 0)</em>\n</p>\n<div class=\"tip tip_4\" id=\"aid4704724_menu\" style=\"display: none\" disautofocus=\"true\">\n<div>\n<p>\n<a href=\"forum.php?mod=attachment&amp;aid=NDcwNDcyNHwzNTU1NTkwYnwxNjgyNzYxNDAwfDB8MTI5MzQ0MQ%3D%3D&amp;nothumb=yes\" target=\"_blank\">下载附件</a>\n\n</p>\n<p>\n<span class=\"y\"><span title=\"2023-04-29 16:25\">1&nbsp;小时前</span> 上传</span>\n<a href=\"javascript:;\" onclick=\"imageRotate('aimg_4704724', 1)\"><img src=\"static/image/common/rleft.gif\" class=\"vm\" /></a>\n<a href=\"javascript:;\" onclick=\"imageRotate('aimg_4704724', 2)\"><img src=\"static/image/common/rright.gif\" class=\"vm\" /></a>\n</p>\n</div>\n<div class=\"tip_horn\"></div>\n</div>\n<p class=\"mbn\">\n\n</p>\n\n\n\n<div class=\"mbn savephotop\">\n\n<img id=\"aimg_4704724\" aid=\"4704724\" src=\"static/image/common/none.gif\" zoomfile=\"https://xksm54s.com/tupian/forum/202304/29/162523uyubk0ykfchkyyuy.jpg\" file=\"https://xksm54s.com/tupian/forum/202304/29/162523uyubk0ykfchkyyuy.jpg\" class=\"zoom\" onclick=\"zoom(this, this.src, 0, 0, 0)\" width=\"600\" alt=\"mmexport1682756678137.jpg\" title=\"mmexport1682756678137.jpg\" w=\"800\" />\n\n</div>\n\n</dd>\n</dl>\n\n</ignore_js_op>\n\n</div>\n\n</div>\n<div id=\"comment_10794090\" class=\"cm\">\n</div>\n\n<div id=\"post_rate_div_10794090\"></div>\n</div>\n</div>\n\n</td></tr>\n<tr><td class=\"plc plm\">\n\n<div id=\"p_btn\" class=\"mtw mbm hm cl\">\n\n<a href=\"home.php?mod=spacecp&amp;ac=favorite&amp;type=thread&amp;id=1293441&amp;formhash=786d433d\" id=\"k_favorite\" onclick=\"showWindow(this.id, this.href, 'get', 0);\" onmouseover=\"this.title = $('favoritenumber').innerHTML + ' 人收藏'\" title=\"收藏本帖\"><i><img src=\"static/image/common/fav.gif\" alt=\"收藏\" />收藏<span id=\"favoritenumber\" style=\"display:none\">0</span></i></a>\n</div>\n\n</td>\n</tr>\n<tr id=\"_postposition10794090\"></tr>\n<tr>\n<td class=\"pls\"></td>\n<td class=\"plc\" style=\"overflow:visible;\">\n<div class=\"po hin\">\n<div class=\"pob cl\">\n<em>\n<a class=\"fastre\" href=\"forum.php?mod=post&amp;action=reply&amp;fid=142&amp;tid=1293441&amp;reppost=10794090&amp;extra=page%3D1&amp;page=1\" onclick=\"showWindow('reply', this.href)\">回复</a>\n</em>\n\n<p>\n<a href=\"javascript:;\" id=\"mgc_post_10794090\" onmouseover=\"showMenu(this.id)\" class=\"showmenu\">使用道具</a>\n<a href=\"javascript:;\" onclick=\"showWindow('miscreport10794090', 'misc.php?mod=report&rtype=post&rid=10794090&tid=1293441&fid=142', 'get', -1);return false;\">举报</a>\n</p>\n\n<ul id=\"mgc_post_10794090_menu\" class=\"p_pop mgcmn\" style=\"display: none;\">\n</ul>\n<script type=\"text/javascript\" reload=\"1\">checkmgcmn('post_10794090')</script>\n</div>\n</div>\n</td>\n</tr>\n<tr class=\"ad\">\n<td class=\"pls\">\n</td>\n<td class=\"plc\">\n</td>\n</tr>\n</table>\n<script type=\"text/javascript\" reload=\"1\">\naimgcount[10794090] = ['4704724'];\nattachimggroup(10794090);\nvar aimgfid = 0;\n</script>\n</div>\n\n<div class=\"show-text2 pad-tb-10\"><a href=\"https://tz112.com\" target=\"_blank\">亚博美女赌场</a>\n<a href=\"https://5656456565.com/vote_topic_80615.do\" target=\"_blank\">澳门现金赌场</a>\n<a href=\"http://23.225.52.51:4466/vip103.html\" target=\"_blank\">澳门新葡京</a>\n<a href=\"https://by301.hdzshbk.cn/?channelCode=by_147\" target=\"_blank\">鲍鱼全球黄播</a>\n<a href=\"http://23.224.188.2:2939/vip107.html\" target=\"_blank\">澳门威尼斯人</a>\n<a href=\"https://154.82.93.23:443/vip058\" target=\"_blank\">澳门葡京赌场</a>\n<a href=\"https://66861056.app/registered\" target=\"_blank\">天天送红包</a>\n<a href=\"https://5581268.cc:8443/?shareName=5581268.cc\" target=\"_blank\">【开元棋牌】</a>\n<a href=\"https://h1116.cc:15555\" target=\"_blank\">现金棋牌</a>\n<a href=\"https://www.rckif.com/001.html\" target=\"_blank\">91成人抖音</a>\n</div>\n<table class=\"plhin\">\n<tr class=\"ad\">\n<td class=\"pls\"></td>\n<td class=\"plc\"></td>\n</tr>\n</table>\n<div class=\"show-text4 pad-tb-10\">\n<div class=\"show-info\">\n<div class=\"avatar\">AD</div>\n</div>\n<div class=\"show-main\"><div><a href=\"https://www.d23dx.com\" class=\"item js-appJump\">\n    <h4>鲍鱼盒子，全网高清直播资源‖自拍视频在线看</h4>\n    <p>注册免费观看，邀请好友送现金红包。</p>\n</a>\n<a href=\"https://www.d23dx.com\" class=\"item js-appJump\">\n    <h4>鲍鱼盒子，全网高清直播资源‖自拍视频在线看</h4>\n    <p>注册免费观看，邀请好友送现金红包。</p>\n</a></div></div>\n</div>\n<table class=\"plhin\">\n<tr class=\"ad\">\n<td class=\"pls\"></td>\n<td class=\"plc\"></td>\n</tr>\n</table>\n<div id=\"postlistreply\" class=\"pl\"><div id=\"post_new\" class=\"viewthread_table\" style=\"display: none\"></div></div>\n</div>\n\n\n<form method=\"post\" autocomplete=\"off\" name=\"modactions\" id=\"modactions\">\n<input type=\"hidden\" name=\"formhash\" value=\"786d433d\" />\n<input type=\"hidden\" name=\"optgroup\" />\n<input type=\"hidden\" name=\"operation\" />\n<input type=\"hidden\" name=\"listextra\" value=\"page%3D1\" />\n<input type=\"hidden\" name=\"page\" value=\"1\" />\n</form>\n\n\n\n<div class=\"pgs mtm mbm cl\">\n<span class=\"pgb y\" id=\"visitedforumstmp\" onmouseover=\"$('visitedforums').id = 'visitedforumstmp';this.id = 'visitedforums';showMenu({'ctrlid':this.id,'pos':'21'})\"><a href=\"forum-142-1.html\">返回列表</a></span>\n<a id=\"newspecialtmp\" onmouseover=\"$('newspecial').id = 'newspecialtmp';this.id = 'newspecial';showMenu({'ctrlid':this.id})\" onclick=\"showWindow('newthread', 'forum.php?mod=post&action=newthread&fid=142')\" href=\"javascript:;\" title=\"发新帖\"><img src=\"static/image/common/pn_post.png\" alt=\"发新帖\" /></a>\n</div>\n\n<!--[diy=diyfastposttop]--><div id=\"diyfastposttop\" class=\"area\"></div><!--[/diy]-->\n<script type=\"text/javascript\">\nvar postminchars = parseInt('10');\nvar postmaxchars = parseInt('500000000');\nvar disablepostctrl = parseInt('0');\n</script>\n\n<div id=\"f_pst\" class=\"pl bm bmw\">\n<form method=\"post\" autocomplete=\"off\" id=\"fastpostform\" action=\"forum.php?mod=post&amp;action=reply&amp;fid=142&amp;tid=1293441&amp;extra=page%3D1&amp;replysubmit=yes&amp;infloat=yes&amp;handlekey=fastpost\" onSubmit=\"return fastpostvalidate(this)\">\n<table cellspacing=\"0\" cellpadding=\"0\">\n<tr>\n<td class=\"pls\">\n</td>\n<td class=\"plc\">\n\n<span id=\"fastpostreturn\"></span>\n\n\n<div class=\"cl\">\n<div id=\"fastsmiliesdiv\" class=\"y\"><div id=\"fastsmiliesdiv_data\"><div id=\"fastsmilies\"></div></div></div><div class=\"hasfsl\" id=\"fastposteditor\">\n<div class=\"tedt mtn\">\n<div class=\"bar\">\n<span class=\"y\">\n<a href=\"forum.php?mod=post&amp;action=reply&amp;fid=142&amp;tid=1293441\" onclick=\"return switchAdvanceMode(this.href)\">高级模式</a>\n</span><script src=\"static/js/seditor.js?Um1\" type=\"text/javascript\"></script>\n<div class=\"fpd\">\n<a href=\"javascript:;\" title=\"文字加粗\" class=\"fbld\">B</a>\n<a href=\"javascript:;\" title=\"设置文字颜色\" class=\"fclr\" id=\"fastpostforecolor\">Color</a>\n<a id=\"fastpostimg\" href=\"javascript:;\" title=\"图片\" class=\"fmg\">Image</a>\n<a id=\"fastposturl\" href=\"javascript:;\" title=\"添加链接\" class=\"flnk\">Link</a>\n<a id=\"fastpostquote\" href=\"javascript:;\" title=\"引用\" class=\"fqt\">Quote</a>\n<a id=\"fastpostcode\" href=\"javascript:;\" title=\"代码\" class=\"fcd\">Code</a>\n<a href=\"javascript:;\" class=\"fsml\" id=\"fastpostsml\">Smilies</a>\n</div></div>\n<div class=\"area\">\n<div class=\"pt hm\">\n您需要登录后才可以回帖 <a href=\"member.php?mod=logging&amp;action=login\" onclick=\"showWindow('login', this.href)\" class=\"xi2\">登录</a> | <a href=\"member.php?mod=register\" class=\"xi2\">立即注册</a>\n</div>\n</div>\n</div>\n</div>\n</div>\n<div id=\"seccheck_fastpost\">\n</div>\n\n\n<input type=\"hidden\" name=\"formhash\" value=\"786d433d\" />\n<input type=\"hidden\" name=\"usesig\" value=\"\" />\n<input type=\"hidden\" name=\"subject\" value=\"  \" />\n<p class=\"ptm pnpost\">\n<a href=\"home.php?mod=spacecp&amp;ac=credit&amp;op=rule&amp;fid=142\" class=\"y\" target=\"_blank\">本版积分规则</a>\n<button type=\"button\" onclick=\"showWindow('login', 'member.php?mod=logging&action=login&guestmessage=yes')\"  onmouseover=\"checkpostrule('seccheck_fastpost', 'ac=reply');this.onmouseover=null\" name=\"replysubmit\" id=\"fastpostsubmit\" class=\"pn pnc vm\" value=\"replysubmit\" tabindex=\"5\"><strong>发表回复</strong></button>\n<label for=\"fastpostrefresh\"><input id=\"fastpostrefresh\" type=\"checkbox\" class=\"pc\" />回帖后跳转到最后一页</label>\n<script type=\"text/javascript\">if(getcookie('fastpostrefresh') == 1) {$('fastpostrefresh').checked=true;}</script>\n</p>\n</td>\n</tr>\n</table>\n</form>\n</div>\n\n<div id=\"visitedforums_menu\" class=\"p_pop blk cl\" style=\"display: none;\">\n<table cellspacing=\"0\" cellpadding=\"0\">\n<tr>\n<td id=\"v_forums\">\n<h3 class=\"mbn pbn bbda xg1\">浏览过的版块</h3>\n<ul class=\"xl xl1\">\n<li><a href=\"forum-36-1.html\">亚洲无码原创</a></li><li><a href=\"forum-2-1.html\">国产原创</a></li></ul>\n</td>\n</tr>\n</table>\n</div>\n<script type=\"text/javascript\">\nnew lazyload();\n</script>\n<script type=\"text/javascript\">document.onkeyup = function(e){keyPageScroll(e, 0, 0, 'forum.php?mod=viewthread&tid=1293441', 1);}</script>\n</div>\n\n<div class=\"wp mtn\">\n<!--[diy=diy3]--><div id=\"diy3\" class=\"area\"></div><!--[/diy]-->\n</div>\n\n<script type=\"text/javascript\">\nfunction succeedhandle_followmod(url, msg, values) {\nvar fObj = $('followmod_'+values['fuid']);\nif(values['type'] == 'add') {\nfObj.innerHTML = '不收听';\nfObj.href = 'home.php?mod=spacecp&ac=follow&op=del&fuid='+values['fuid'];\n} else if(values['type'] == 'del') {\nfObj.innerHTML = '收听TA';\nfObj.href = 'home.php?mod=spacecp&ac=follow&op=add&hash=786d433d&fuid='+values['fuid'];\n}\n}\nfixed_avatar([10794090], 1);\n</script><div class=\"show-text cl\"><a href=\"https://tz112.com\" target=\"_blank\" class=\"js-randomBg\">亚博美女赌场</a>\n<a href=\"https://5656456565.com/vote_topic_80615.do\" target=\"_blank\" class=\"js-randomBg\">澳门现金赌场</a>\n<a href=\"http://23.225.52.51:4466/vip103.html\" target=\"_blank\" class=\"js-randomBg\">澳门新葡京</a>\n<a href=\"https://by301.hdzshbk.cn/?channelCode=by_147\" target=\"_blank\" class=\"js-randomBg\">鲍鱼全球黄播</a>\n<a href=\"http://23.224.188.2:2939/vip107.html\" target=\"_blank\" class=\"js-randomBg\">澳门威尼斯人</a>\n<a href=\"https://154.82.93.23:443/vip058\" target=\"_blank\" class=\"js-randomBg\">澳门葡京赌场</a>\n<a href=\"https://66861056.app/registered\" target=\"_blank\" class=\"js-randomBg\">天天送红包</a>\n<a href=\"https://5581268.cc:8443/?shareName=5581268.cc\" target=\"_blank\" class=\"js-randomBg\">【开元棋牌】</a>\n<a href=\"https://h1116.cc:15555\" target=\"_blank\" class=\"js-randomBg\">现金棋牌</a>\n<a href=\"https://www.rckif.com/001.html\" target=\"_blank\" class=\"js-randomBg\">91成人抖音</a>\n</div>\n<div class=\"h10\"></div>\n\t</div>\n<div id=\"ft\" class=\"wp cl\">\n<div id=\"flk\" class=\"y\">\n<p>\n广告联系：sehuatang@gmail.com<span class=\"pipe\">|</span>\n<a href=\"forum.php?mod=misc&action=showdarkroom\" >小黑屋</a><span class=\"pipe\">|</span><strong><a href=\"https://www.sehuatang.net\" target=\"_blank\">98堂[原色花堂]</a></strong>\n</p>\n<p class=\"xs0\">\nGMT+8, 2023-04-29 17:43<span id=\"debuginfo\">\n, Processed in 0.036039 second(s), Total 4, Slave 3 queries\n, Redis On.\n</span>\n</p>\n</div>\n<div id=\"frt\">\n<p>Powered by <strong><a href=\"http://www.discuz.net\" target=\"_blank\">Discuz!</a></strong> <em>X3.4</em></p>\n<p class=\"xs0\">Copyright &copy; 2001-2021, Tencent Cloud.</p>\n</div></div>\n\n<div id=\"scrolltop\">\n<span><a href=\"forum.php?mod=post&amp;action=reply&amp;fid=142&amp;tid=1293441&amp;extra=page%3D1&amp;page=1\" onclick=\"showWindow('reply', this.href)\" class=\"replyfast\" title=\"快速回复\"><b>快速回复</b></a></span>\n<span hidefocus=\"true\"><a title=\"返回顶部\" onclick=\"window.scrollTo('0','0')\" class=\"scrolltopa\" ><b>返回顶部</b></a></span>\n<span>\n<a href=\"forum-142-1.html\" hidefocus=\"true\" class=\"returnlist\" title=\"返回列表\"><b>返回列表</b></a>\n</span>\n</div>\n<script type=\"text/javascript\">_attachEvent(window, 'scroll', function () { showTopLink(); });checkBlind();randomBg('js-randomBg');</script>\n</body>\n</html>\n"
	config := "{    \"name_content\": \"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/table[1]/tbody/tr/td[@class='plc ptm pbn vwthd']/h1[@class='ts']/span[@id='thread_subject']\",\n\"taq_content\": \"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/table[1]/tbody/tr/td[@class='plc ptm pbn vwthd']/h1[@class='ts']/a\",\n\"pubishTime_content\": \"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/div[@id]/table[@id]/tbody/tr[1]/td[@class='plc']/div[@class='pi']/div[@class='pti']/div[@class='authi']/em[@id]\",\n\"start_num_content\": \"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/div[@id]/table[@id]/tbody/tr[2]/td[@class='plc plm']/div[@id='p_btn']/a[@id='k_favorite']/i/span[@id='favoritenumber']\",\n\"discussNum_content\": \"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/table[1]/tbody/tr/td[@class='pls ptn pbn']/div[@class='hm ptn']/span[@class='xi1'][2]\",\n\"pv_content\": \"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/table[1]/tbody/tr/td[@class='pls ptn pbn']/div[@class='hm ptn']/span[@class='xi1'][1]\",\n\"evaluate_content\": \"/html/body[@id='nv_forum']/div[@id='wp']/div[@id='ct']/div[@id='postlist']/div[@id]/table[@id]/tbody/tr[1]/td[@class='plc']/div[@class='pct']/div[@class='pcb']/dl[@id]/dd/table[@class='ratl']/tbody[1]/tr/th[@class='xw1']\"\n}"
	srt := ""
	fmt.Println(Exe(config, srt, html, ""))
}

func TestMap(t *testing.T) {
	config := "{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\",\"obj\":{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\"}}"
	configJson := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(config), &configJson) //第二个参数要地址传递
	if err != nil {
		printlnErr(err, "Parser config err")
		return
	}
	fmt.Println(configJson)
}
