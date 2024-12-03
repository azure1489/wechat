package message

import "encoding/xml"

// <msg>
//
//	<appmsg appid="" sdkver="0">
//	  <title><![CDATA[活动进度通知]]></title>
//	  <des><![CDATA[活动名称:助力膨胀
//
// 活动进度:马上成功了，快去邀请好友吧！
// 温馨提示:点击查看活动详情>>
// ]]></des>
//
//	<action/>
//	<type>5</type>
//	<showtype>1</showtype>
//	<soundtype>0</soundtype>
//	<content><![CDATA[]]></content>
//	<contentattr>0</contentattr>
//	<url><![CDATA[outer_packages/mactivitydyna/pages/boost/boost.html?activityKey=43ea6d692ff44ee98ef52cf9ff59f72c]]></url>
//	<lowurl><![CDATA[]]></lowurl>
//	<appattach>
//	  <totallen>0</totallen>
//	  <attachid/>
//	  <fileext/>
//	  <cdnthumburl><![CDATA[]]></cdnthumburl>
//	  <cdnthumbaeskey><![CDATA[]]></cdnthumbaeskey>
//	  <aeskey><![CDATA[]]></aeskey>
//	</appattach>
//	<extinfo/>
//	<sourceusername><![CDATA[]]></sourceusername>
//	<sourcedisplayname><![CDATA[]]></sourcedisplayname>
//	<mmreader>
//	  <category type="0" count="1">
//	    <name><![CDATA[惠省 I 吃喝玩乐小助手]]></name>
//	    <topnew>
//	      <cover><![CDATA[]]></cover>
//	      <width>0</width>
//	      <height>0</height>
//	      <digest><![CDATA[活动名称:助力膨胀
//
// 活动进度:马上成功了，快去邀请好友吧！
// 温馨提示:点击查看活动详情>>
// ]]></digest>
//
//	</topnew>
//	<item>
//	  <itemshowtype>4</itemshowtype>
//	  <title><![CDATA[活动进度通知]]></title>
//	  <url><![CDATA[outer_packages/mactivitydyna/pages/boost/boost.html?activityKey=43ea6d692ff44ee98ef52cf9ff59f72c]]></url>
//	  <shorturl><![CDATA[]]></shorturl>
//	  <longurl><![CDATA[]]></longurl>
//	  <pub_time>1733230847</pub_time>
//	  <cover><![CDATA[]]></cover>
//	  <tweetid/>
//	  <digest><![CDATA[活动名称:助力膨胀
//
// 活动进度:马上成功了，快去邀请好友吧！
// 温馨提示:点击查看活动详情>>
// ]]></digest>
//
//	        <fileid>0</fileid>
//	        <sources>
//	          <source>
//	            <name><![CDATA[惠省 I 吃喝玩乐小助手]]></name>
//	          </source>
//	        </sources>
//	        <styles>
//	          <topColor><![CDATA[]]></topColor>
//	        </styles>
//	        <native_url/>
//	        <del_flag>0</del_flag>
//	        <contentattr>0</contentattr>
//	        <play_length>0</play_length>
//	        <play_url><![CDATA[]]></play_url>
//	        <voice_id><![CDATA[]]></voice_id>
//	        <tid><![CDATA[]]></tid>
//	        <nonce_id><![CDATA[]]></nonce_id>
//	        <voice_type>0</voice_type>
//	        <player><![CDATA[]]></player>
//	        <template_op_type>1</template_op_type>
//	        <weapp_username><![CDATA[gh_9dda55bf7807@app]]></weapp_username>
//	        <weapp_path><![CDATA[outer_packages/mactivitydyna/pages/boost/boost.html?activityKey=43ea6d692ff44ee98ef52cf9ff59f72c]]></weapp_path>
//	        <weapp_version>0</weapp_version>
//	        <weapp_state>0</weapp_state>
//	        <music_source>0</music_source>
//	        <pic_num>0</pic_num>
//	        <show_complaint_button>0</show_complaint_button>
//	        <vid><![CDATA[]]></vid>
//	        <recommendation><![CDATA[]]></recommendation>
//	        <pic_urls/>
//	        <multi_picture_cover/>
//	        <comment_topic_id>0</comment_topic_id>
//	        <cover_235_1><![CDATA[]]></cover_235_1>
//	        <cover_1_1><![CDATA[]]></cover_1_1>
//	        <cover_16_9><![CDATA[]]></cover_16_9>
//	        <appmsg_like_type>0</appmsg_like_type>
//	        <video_width>0</video_width>
//	        <video_height>0</video_height>
//	        <is_pay_subscribe>0</is_pay_subscribe>
//	        <summary><![CDATA[]]></summary>
//	        <general_string><![CDATA[]]></general_string>
//	        <finder_feed/>
//	        <finder_live/>
//	        <right_cover_url/>
//	        <text_title><![CDATA[]]></text_title>
//	        <has_redpacket_cover>0</has_redpacket_cover>
//	        <title_v2><![CDATA[]]></title_v2>
//	        <product_activity/>
//	      </item>
//	    </category>
//	    <publisher>
//	      <username><![CDATA[gh_9dda55bf7807@app]]></username>
//	      <nickname><![CDATA[惠省 I 吃喝玩乐小助手]]></nickname>
//	    </publisher>
//	    <template_header>
//	      <title><![CDATA[活动进度通知]]></title>
//	      <title_color><![CDATA[]]></title_color>
//	      <pub_time>1733230847</pub_time>
//	      <first_data><![CDATA[]]></first_data>
//	      <first_color><![CDATA[]]></first_color>
//	    </template_header>
//	    <template_detail>
//	      <template_show_type>1</template_show_type>
//	      <text_content>
//	        <cover><![CDATA[]]></cover>
//	        <text><![CDATA[]]></text>
//	        <color><![CDATA[]]></color>
//	      </text_content>
//	      <line_content>
//	        <lines>
//	          <line>
//	            <key>
//	              <word><![CDATA[活动名称]]></word>
//	              <color><![CDATA[#888888]]></color>
//	            </key>
//	            <value>
//	              <word><![CDATA[助力膨胀]]></word>
//	              <color><![CDATA[#000000]]></color>
//	            </value>
//	          </line>
//	          <line>
//	            <key>
//	              <word><![CDATA[活动进度]]></word>
//	              <color><![CDATA[#888888]]></color>
//	            </key>
//	            <value>
//	              <word><![CDATA[马上成功了，快去邀请好友吧！]]></word>
//	              <color><![CDATA[#000000]]></color>
//	            </value>
//	          </line>
//	          <line>
//	            <key>
//	              <word><![CDATA[温馨提示]]></word>
//	              <color><![CDATA[#888888]]></color>
//	            </key>
//	            <value>
//	              <word><![CDATA[点击查看活动详情>>]]></word>
//	              <color><![CDATA[#000000]]></color>
//	            </value>
//	          </line>
//	        </lines>
//	      </line_content>
//	      <opitems>
//	        <opitem>
//	          <word><![CDATA[进入小程序查看]]></word>
//	          <url><![CDATA[outer_packages/mactivitydyna/pages/boost/boost.html?activityKey=43ea6d692ff44ee98ef52cf9ff59f72c]]></url>
//	          <icon><![CDATA[]]></icon>
//	          <color><![CDATA[#000000]]></color>
//	          <weapp_username><![CDATA[gh_9dda55bf7807@app]]></weapp_username>
//	          <weapp_path><![CDATA[outer_packages/mactivitydyna/pages/boost/boost.html?activityKey=43ea6d692ff44ee98ef52cf9ff59f72c]]></weapp_path>
//	          <op_type>1</op_type>
//	          <weapp_version>0</weapp_version>
//	          <weapp_state>0</weapp_state>
//	          <hint_word><![CDATA[]]></hint_word>
//	          <is_rich_text>0</is_rich_text>
//	          <display_line_number>0</display_line_number>
//	          <general_string><![CDATA[]]></general_string>
//	          <is_show_red_dot>0</is_show_red_dot>
//	          <ext_id><![CDATA[]]></ext_id>
//	          <business_id><![CDATA[]]></business_id>
//	          <thumbnail><![CDATA[]]></thumbnail>
//	          <is_show_play_btn>0</is_show_play_btn>
//	          <dmicon><![CDATA[]]></dmicon>
//	        </opitem>
//	        <show_type>0</show_type>
//	      </opitems>
//	      <template_ext>
//	        <we_app_state>0</we_app_state>
//	        <we_app_version>0</we_app_version>
//	        <we_app_latest_version>103</we_app_latest_version>
//	        <is_audio_template>0</is_audio_template>
//	      </template_ext>
//	      <new_tmpl_type>2</new_tmpl_type>
//	      <flat_content>
//	        <lines>
//	          <line>
//	            <key>
//	              <word><![CDATA[活动名称]]></word>
//	              <color><![CDATA[#888888]]></color>
//	            </key>
//	            <value>
//	              <word><![CDATA[助力膨胀]]></word>
//	              <color><![CDATA[#000000]]></color>
//	            </value>
//	          </line>
//	          <line>
//	            <key>
//	              <word><![CDATA[活动进度]]></word>
//	              <color><![CDATA[#888888]]></color>
//	            </key>
//	            <value>
//	              <word><![CDATA[马上成功了，快去邀请好友吧！]]></word>
//	              <color><![CDATA[#000000]]></color>
//	            </value>
//	          </line>
//	          <line>
//	            <key>
//	              <word><![CDATA[温馨提示]]></word>
//	              <color><![CDATA[#888888]]></color>
//	            </key>
//	            <value>
//	              <word><![CDATA[点击查看活动详情>>]]></word>
//	              <color><![CDATA[#000000]]></color>
//	            </value>
//	          </line>
//	        </lines>
//	      </flat_content>
//	    </template_detail>
//	    <forbid_forward>0</forbid_forward>
//	    <notify_msg/>
//	  </mmreader>
//	  <thumburl><![CDATA[]]></thumburl>
//	  <template_id><![CDATA[6MNT7TDluRMtMuYCVQ8zs22k4Vi1gyBfjwsMPdlwIU4]]></template_id>
//	</appmsg>
//	<fromusername><![CDATA[gh_9dda55bf7807@app]]></fromusername>
//	<appinfo>
//	  <version>103</version>
//	  <appname><![CDATA[惠省 I 吃喝玩乐小助手]]></appname>
//	  <isforceupdate>1</isforceupdate>
//	</appinfo>
//
// </msg>
type ServiceNoticeXml struct {
	XMLName xml.Name `xml:"msg"`
	AppMsg  struct {
		AppID    string `xml:"appid,attr"`
		SDKVer   string `xml:"sdkver,attr"`
		Title    string `xml:"title"`
		Des      string `xml:"des"`
		Type     int    `xml:"type"`
		URL      string `xml:"url"`
		MMReader struct {
			Category struct {
				Type  int    `xml:"type,attr"`
				Count int    `xml:"count,attr"`
				Name  string `xml:"name"`
				Item  struct {
					Title         string `xml:"title"`
					URL           string `xml:"url"`
					WeappUsername string `xml:"weapp_username"`
					WeappPath     string `xml:"weapp_path"`
				} `xml:"item"`
			} `xml:"category"`
			Publisher struct {
				Username string `xml:"username"`
				Nickname string `xml:"nickname"`
			} `xml:"publisher"`
		} `xml:"mmreader"`
	} `xml:"appmsg"`
	FromUsername string `xml:"fromusername"`
	AppInfo      struct {
		Version       string `xml:"version"`
		AppName       string `xml:"appname"`
		IsForceUpdate int    `xml:"isforceupdate"`
	} `xml:"appinfo"`
}

type ServiceNotice struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	AppName     string `json:"appName"`
	WeappPath   string `json:"weappPath"`
	WeappUser   string `json:"weappUser"`
}
