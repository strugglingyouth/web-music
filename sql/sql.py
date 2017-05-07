{u'album': {u'picUrl': u'http://p1.music.126.net/vVmZs-NuZBMPrQxn-doG-w==/5715261441290700.jpg', u'pic': 5715261441290700, u'id': 2685023, u'tns': [], u'name': u'\u7ec8\u4e8e\u7b49\u5230\u4f60'}, u'status': 0, u'commentThreadId': u'R_SO_4_27836179', u'name': u'\u7ec8\u4e8e\u7b49\u5230\u4f60', u'djid': 0, u'fee': 0, u'duration': 295214, u'no': 0, u'mvid': 193171, u'alias': [], u'score': 100.0, u'ftype': 0, u'artists': [{u'alias': [], u'id': 10561, u'tns': [], u'name': u'\u5f20\u9753\u9896'}], u'transNames': None, u'copyrightId': 0, u'type': 0, u'id': 27836179, u'privilege': {u'toast': False, u'dl': 320000, u'fee': 0, u'subp': 1, u'sp': 7, u'st': 0, u'payed': 0, u'flag': 0, u'maxbr': 999000, u'cs': False, u'cp': 1, u'fl': 320000, u'id': 27836179, u'pl': 320000}}



3. 获取歌手专辑列表

GET http://music.163.com/api/artist/albums/[artist_id]/

其中artist_id用歌手id替换

参数

offset: 偏移数量，用于分页

limit: 返回数量

示例

curl -b "appver=1.5.2;" "http://music.163.com/api/artist/albums/10557?offset=0&limit=3"



{
    "commentThreadId": "R_SO_4_145586",
    "status": 0,
    "fee": 0,
    "score": 100,
    "no": 1,
    "id": 145586,
    "mvid": 5436513,
    "name": "拯救",         #歌曲名称
    "alias": [
        "电视剧《拿什么拯救你，我的爱人》片尾曲"
    ],
    "duration": 332591,
    "album": {
        "picUrl": "http://p1.music.126.net/IBCYtNlA1-vVpe3mj9F-Jw==/130841883718470.jpg",
        "pic": 130841883718470,
        "id": 14574,                #  歌曲id
        "tns": [],
        "name": "缘份的天空"        #专辑
    },
    "ftype": 0,
    "artists": [
        {
            "alias": [],
            "id": 4941,           # 歌手 id
            "tns": [],
            "name": "孙楠"        #歌手
        }
    ],
    "transNames": null,
    "copyrightId": 14031,
    "type": 0,
    "djid": 0,
    "privilege": {
        "toast": false,
        "dl": 320000,
        "fee": 0,
        "subp": 1,
        "sp": 7,
        "st": 0,
        "payed": 0,
        "flag": 0,
        "maxbr": 320000,
        "cs": false,
        "cp": 1,
        "fl": 320000,
        "id": 145586,
        "pl": 320000
    }
}

抓取的精彩评论 

{
    "liked": false,
    "beReplied": [],
    "content": "还有五十多天高考，还有五十多天就该say good goodbye",
    "user": {
        "remarkName": null,
        "expertTags": null,
        "avatarUrl": "http://p1.music.126.net/7XjYBi60IiPgqXUKV3JEPw==/1378787598224454.jpg",
        "userId": 122765071,
        "locationInfo": null,
        "userType": 0,
        "authStatus": 0,
        "nickname": "_入门唯觉一庭香",
        "vipType": 0
    },
    "likedCount": 3847,
    "time": 1492227459820,
    "commentId": 356932948
}




# 歌手信息表
create table IF NOT EXISTS singer(
singer_id  int(11),               # 歌手id
singer_name  varchar(255)         # 歌手名字
);


# 歌曲信息表
create table IF NOT EXISTS music_items(
music_id  int(11),               #歌曲id
music_url  varchar(255),         # 歌曲 url
music_name  varchar(255),        # 歌曲名
music_album  varchar(255),       # 歌曲专辑
music_mv     varchar(255)        # 歌曲 mv
music_artist varchar(255),       # 歌手，外键关联到歌手信息表
music_artist_id int(11)         # 歌手id ， 外键
);

# 歌曲评论表

create table IF NOT EXISTS musiccomment(
music_id  int(11),       #外键

comment_content  varchar(10000),   #评论内容
comment_user_id  int(11),           #评论用户 id 
comment_username varchar(200),      # 评论用户名
comment_like_count int(11),         # 评论喜欢数
comment_id      int(11)             # 评论 id
);

#音乐榜单
create table IF NOT EXISTS hot_uk(
music_id  int(11),               #歌曲id
music_url  varchar(255),         # 歌曲 url
music_name  varchar(255),        # 歌曲名
music_album  varchar(255),       # 歌曲专辑
music_artist varchar(255),      # 歌手，外键关联到歌手信息表
music_artist_id  int(11)
);




#百度音乐排行  新歌榜(newsong)  热歌榜mod-hotsong  网络歌曲榜mod-netsong   原创榜mod-originsong
create table IF NOT EXISTS baidu_originsong(
music_id  int(11),               #歌曲id
music_url   varchar(255),        # 歌曲链接
music_name  varchar(255),        # 歌曲名

music_artist varchar(255),      # 歌手，外键关联到歌手信息表
music_artist_url varchar(255)
);


<html>
 <head></head>
 <body>
  <div class="opera-icon" data-args="{&quot;id&quot;:&quot;540560826&quot;,&quot;type&quot;:&quot;song&quot;,&quot;moduleName&quot;:&quot;newIcon&quot;,&quot;albumId&quot;:&quot;540560824&quot;,&quot;resourceTypeExt&quot;:&quot;0&quot;,&quot;siPresaleFlag&quot;:null,&quot;mediaType&quot;:1,&quot;songPic&quot;:&quot;http:\/\/musicdata.baidu.com\/data2\/pic\/ab055cc7c59de25a37f06f060ec3bccf\/540561450\/540561450.jpg@s_0,w_90&quot;,&quot;songTitle&quot;:&quot;\u6211\u5bb3\u6015&quot;,&quot;songPublishTime&quot;:null}">
   <a href="javascript:;" class="opera-icon-play icon icon-music-play js-play-song"></a>
   <a href="javascript:;" class="opera-icon-add icon icon-music-collect js-add"></a>
  </div>
  <li class=" top2">
   <div class="index">
    02
   </div>
   <div class="status status-steady">
    <i class="icon-status"></i>
   </div>
   <div class="song-info">
    <div class="info">
     <div class="song">
      <a href="/song/540175998" title="暧昧">暧昧</a>
      <span class="artist"> <span class="author_list" title="薛之谦"> <a hidefocus="true" href="/artist/2517">薛之谦</a> </span></span>
     </div>
    </div>
   </div></li>
 </body>
</html>