{u'album': {u'picUrl': u'http://p1.music.126.net/vVmZs-NuZBMPrQxn-doG-w==/5715261441290700.jpg', u'pic': 5715261441290700, u'id': 2685023, u'tns': [], u'name': u'\u7ec8\u4e8e\u7b49\u5230\u4f60'}, u'status': 0, u'commentThreadId': u'R_SO_4_27836179', u'name': u'\u7ec8\u4e8e\u7b49\u5230\u4f60', u'djid': 0, u'fee': 0, u'duration': 295214, u'no': 0, u'mvid': 193171, u'alias': [], u'score': 100.0, u'ftype': 0, u'artists': [{u'alias': [], u'id': 10561, u'tns': [], u'name': u'\u5f20\u9753\u9896'}], u'transNames': None, u'copyrightId': 0, u'type': 0, u'id': 27836179, u'privilege': {u'toast': False, u'dl': 320000, u'fee': 0, u'subp': 1, u'sp': 7, u'st': 0, u'payed': 0, u'flag': 0, u'maxbr': 999000, u'cs': False, u'cp': 1, u'fl': 320000, u'id': 27836179, u'pl': 320000}}



3. ��ȡ����ר���б�

GET http://music.163.com/api/artist/albums/[artist_id]/

����artist_id�ø���id�滻

����

offset: ƫ�����������ڷ�ҳ

limit: ��������

ʾ��

curl -b "appver=1.5.2;" "http://music.163.com/api/artist/albums/10557?offset=0&limit=3"



{
    "commentThreadId": "R_SO_4_145586",
    "status": 0,
    "fee": 0,
    "score": 100,
    "no": 1,
    "id": 145586,
    "mvid": 5436513,
    "name": "����",         #��������
    "alias": [
        "���Ӿ硶��ʲô�����㣬�ҵİ��ˡ�Ƭβ��"
    ],
    "duration": 332591,
    "album": {
        "picUrl": "http://p1.music.126.net/IBCYtNlA1-vVpe3mj9F-Jw==/130841883718470.jpg",
        "pic": 130841883718470,
        "id": 14574,                #  ����id
        "tns": [],
        "name": "Ե�ݵ����"        #ר��
    },
    "ftype": 0,
    "artists": [
        {
            "alias": [],
            "id": 4941,           # ���� id
            "tns": [],
            "name": "���"        #����
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

ץȡ�ľ������� 

{
    "liked": false,
    "beReplied": [],
    "content": "������ʮ����߿���������ʮ����͸�say good goodbye",
    "user": {
        "remarkName": null,
        "expertTags": null,
        "avatarUrl": "http://p1.music.126.net/7XjYBi60IiPgqXUKV3JEPw==/1378787598224454.jpg",
        "userId": 122765071,
        "locationInfo": null,
        "userType": 0,
        "authStatus": 0,
        "nickname": "_����Ψ��һͥ��",
        "vipType": 0
    },
    "likedCount": 3847,
    "time": 1492227459820,
    "commentId": 356932948
}




# ������Ϣ��
create table IF NOT EXISTS singer(
singer_id  int(11),               # ����id
singer_name  varchar(255)         # ��������
);


# ������Ϣ��
create table IF NOT EXISTS music_items(
music_id  int(11),               #����id
music_url  varchar(255),         # ���� url
music_name  varchar(255),        # ������
music_album  varchar(255),       # ����ר��
music_mv     varchar(255)        # ���� mv
music_artist varchar(255),       # ���֣����������������Ϣ��
music_artist_id int(11)         # ����id �� ���
);

# �������۱�

create table IF NOT EXISTS musiccomment(
music_id  int(11),       #���

comment_content  varchar(10000),   #��������
comment_user_id  int(11),           #�����û� id 
comment_username varchar(200),      # �����û���
comment_like_count int(11),         # ����ϲ����
comment_id      int(11)             # ���� id
);

#���ְ�
create table IF NOT EXISTS hot_uk(
music_id  int(11),               #����id
music_url  varchar(255),         # ���� url
music_name  varchar(255),        # ������
music_album  varchar(255),       # ����ר��
music_artist varchar(255),      # ���֣����������������Ϣ��
music_artist_id  int(11)
);




#�ٶ���������  �¸��(newsong)  �ȸ��mod-hotsong  ���������mod-netsong   ԭ����mod-originsong
create table IF NOT EXISTS baidu_originsong(
music_id  int(11),               #����id
music_url   varchar(255),        # ��������
music_name  varchar(255),        # ������

music_artist varchar(255),      # ���֣����������������Ϣ��
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
      <a href="/song/540175998" title="����">����</a>
      <span class="artist"> <span class="author_list" title="Ѧ֮ǫ"> <a hidefocus="true" href="/artist/2517">Ѧ֮ǫ</a> </span></span>
     </div>
    </div>
   </div></li>
 </body>
</html>