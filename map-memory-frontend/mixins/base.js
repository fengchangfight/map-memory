import { AXIOS } from '~/common/http-commons'

export default {
  methods: {
    getBaseServiceUrl() {
      AXIOS.get('/api/v1/base-service/url').then(response => {
        // JSON responses are automatically parsed.
        this.base_service_url = response.data;
      }).catch(e => {

      })
    },
    getBaseUrl() {
      AXIOS.get('/api/v1/base-url/url').then(response => {
        // JSON responses are automatically parsed.
        this.base_url = response.data;
      }).catch(e => {

      })
    },
    checklogin() {
      AXIOS.get('/api/v1/user/whoami').then(response => {
        // JSON responses are automatically parsed.

        this.isLoggedIn = true;
      }).catch(e => {
        this.isLoggedIn = false;
        this.$router.push('/login')
      })
    },
    goPage(val) {
      this.$router.push(val)
    },
  },
  data() {
    return {
      base_url: '',
      base_service_url: '',
      isLoggedIn: false,
      available_icons: [
        {
          id: 'chrome.png',
          name: 'Chrome'
        }, {
          id: 'firefox.png',
          name: 'Firefox'
        },
        {
          id: 'nvbing.png',
          name: '女兵'
        },
        {
          id: 'zombie.png',
          name: '丧尸'
        },
        {
          id: 'menggudao.png',
          name: '蒙古刀'
        },
        {
          id: 'sword.png',
          name: '剑'
        },
        {
          id: 'bow.png',
          name: '弓箭'
        },
        {
          id: 'dadao.png',
          name: '大刀'
        }, {
          id: 'hongyingqiang.png',
          name: '红缨枪'
        },
        {
          id: 'sanjiegun.png',
          name: '三节棍'
        },
        {
          id: 'knife.png',
          name: '匕首'
        },
        {
          id: 'm.png',
          name: 'M'
        }, {
          id: 'c.png',
          name: 'C'
        },
        {
          id: 'wugui.png',
          name: '乌龟'
        },
        {
          id: 'crocodile.png',
          name: '鳄鱼'
        },
        {
          id: 'yudi.png',
          name: '玉帝'
        },
        {
          id: 'qinshihuang.png',
          name: '秦始皇'
        },
        {
          id: 'liao.png',
          name: '李敖'
        },
        {
          id: 'maozedong.png',
          name: '毛泽东'
        },
        {
          id: 'jetli.png',
          name: '李连杰'
        },
        {
          id: 'nietzsche.png',
          name: '尼采'
        },
        {
          id: 'lixiaolong.png',
          name: '李小龙'
        },
        {
          id: 'plane.png',
          name: '飞机'
        },
        {
          id: 'gaotie.png',
          name: '高铁'
        },
        {
          id: 'car.png',
          name: '小汽车'
        },
        {
          id: 'doctorcap.png',
          name: '博士帽'
        },
        {
          id: 'cup.png',
          name: '茶杯'
        },
        {
          id: 'grenade.png',
          name: '手雷'
        },
        {
          id: 'daodan.png',
          name: '导弹'
        },
        {
          id: 'handgun.png',
          name: '手枪'
        },
        {
          id: 'tank.png',
          name: '坦克'
        },
        {
          id: 'watch.png',
          name: '钟表'
        },
        {
          id: 'chenguohan.png',
          name: '陈国汉'
        },
        {
          id: 'yadianna.png',
          name: '雅典娜'
        },
        {
          id: 'damenwulang.png',
          name: '大门五郎'
        },
        {
          id: 'dongzhang.png',
          name: '东丈'
        },
        {
          id: 'bashenan.png',
          name: '八神庵'
        },
        {
          id: 'wuseng.png',
          name: '武僧'
        },
        {
          id: 'jinguanzhang.png',
          name: '金馆长'
        },
        {
          id: 'yaomingbiaoqingbao.png',
          name: '姚明脸'
        },
        {
          id: 'zhangxueyoubiaoqingbao.png',
          name: '张学友表情'
        },
        {
          id: 'jiaohuang.png',
          name: '教皇表情包表情'
        },
        {
          id: 'nikeladaoba.png',
          name: '你可拉倒吧'
        },
        {
          id: 'samoye.png',
          name: '萨摩耶'
        },
        {
          id: 'chaiquan.png',
          name: '柴犬'
        },
        {
          id: 'chaiquan2.png',
          name: '柴犬2'
        },
        {
          id: 'hashiqi.png',
          name: '哈士奇'
        },
        {
          id: 'hashiqi2.png',
          name: '哈士奇2'
        },
        {
          id: 'hashiqi3.png',
          name: '哈士奇3'
        },
        {
          id: 'alasijiabaobao.png',
          name: '阿拉斯加宝宝'
        },
        {
          id: 'hashiqibaobao.png',
          name: '哈士奇宝宝'
        },
        {
          id: 'keji.png',
          name: '柯基'
        },
        {
          id: 'kejibaobao.png',
          name: '柯基宝宝'
        },
        {
          id: 'bluecat.png',
          name: '蓝猫'
        },
        {
          id: 'persiancat.png',
          name: '波斯猫'
        },
        {
          id: 'garfield.png',
          name: '加菲猫'
        },
        {
          id: 'ragdoll.png',
          name: '布偶猫'
        },
        {
          id: 'horse.png',
          name: '马'
        },
        {
          id: 'penguin.png',
          name: '企鹅'
        },
        {
          id: 'elephant.png',
          name: '大象'
        },
        {
          id: 'rabbit.png',
          name: '兔子'
        },
        {
          id: 'pig.png',
          name: '猪'
        }, {
          id: 'laohu.png',
          name: '老虎'
        }, {
          id: 'shizi.png',
          name: '狮子'
        }, {
          id: 'go.png',
          name: '土拨鼠'
        },
        {
          id: 'hammer.png',
          name: '锤子'
        },
        {
          id: 'shovel.png',
          name: '铲子'
        }, {
          id: 'laohuqian.png',
          name: '老虎钳'
        },
        {
          id: 'sickle.png',
          name: '镰刀'
        },
        {
          id: 'banshou.png',
          name: '扳手'
        },
        {
          id: 'baby.png',
          name: '婴儿'
        },
        {
          id: 'book2.png',
          name: '大金书'
        }, {
          id: 'book5.png',
          name: '书5'
        }, {
          id: 'writing.png',
          name: '写作'
        }, {
          id: 'memo.png',
          name: '便笺'
        }, {
          id: 'book.png',
          name: '秘籍'
        }, {
          id: 'book3.png',
          name: '小蓝书'
        }, {
          id: 'medical.png',
          name: '医药'
        }, {
          id: 'pill.png',
          name: '药丸'
        }, {
          id: 'cyclopedia.png',
          name: '百科'
        }, {
          id: 'food.png',
          name: '食物'
        }, {
          id: 'website.png',
          name: '网站'
        }, {
          id: 'lock.png',
          name: '锁'
        }, {
          id: 'cloud.png',
          name: '云'
        }, {
          id: 'finance.png',
          name: '金融'
        }, {
          id: 'geometry.png',
          name: '图形'
        }, {
          id: 'dialog.png',
          name: '对话'
        }, {
          id: 'spring.png',
          name: 'Spring'
        }, {
          id: 'leetcode.png',
          name: 'Leetcode'
        }, {
          id: 'youtube.png',
          name: 'Youtube'
        }, {
          id: 'mail.png',
          name: '邮件'
        }, {
          id: 'atom.png',
          name: 'atom'
        }, {
          id: 'git.png',
          name: 'git'
        }, {
          id: 'practice.png',
          name: '实践'
        }, {
          id: 'knowledge.png',
          name: '知识'
        }, {
          id: 'redpacket.png',
          name: '红包'
        }, {
          id: 'class.png',
          name: '课程'
        }, {
          id: 'opensource.png',
          name: '开源'
        }, {
          id: 'torch.png',
          name: 'Torch'
        },
        {
          id: 'java.png',
          name: 'Java'
        }, {
          id: 'location.png',
          name: '位置'
        },
        {
          id: 't.png',
          name: 'T'
        }, {
          id: 'django.png',
          name: 'Django'
        }, {
          id: 'darkbook.png',
          name: '暗黑本'
        },
        {
          id: 'vscode.png',
          name: 'vscode'
        },
        {
          id: 'video.png',
          name: '视频'
        },
        {
          id: 'ppt.png',
          name: 'PPT'
        }, {
          id: 'word.png',
          name: 'Word'
        },
        {
          id: 'excel.png',
          name: 'Excel'
        }, {
          id: 'diary.png',
          name: '日记'
        }, {
          id: 'dollar.png',
          name: '美元'
        }, {
          id: 'intellij.png',
          name: 'Intellij'
        }, {
          id: 'eclipse.png',
          name: 'eclipse'
        }, {
          id: 'gradle.png',
          name: 'gradle'
        },
        {
          id: 'book4.png',
          name: '宝书'
        }, {
          id: 'house.png',
          name: '房子'
        }, {
          id: 'premiere.png',
          name: 'Premiere'
        }, {
          id: 'newspaper.png',
          name: '报纸'
        }, {
          id: 'wordpress.png',
          name: 'wordpress'
        }, {
          id: 'heart.png',
          name: '心'
        }, {
          id: 'mofang.png',
          name: '魔方'
        },
        {
          id: 'taiqiu.png',
          name: '台球'
        },
        {
          id: 'tv.png',
          name: '电视'
        },

        {
          id: 'chenglong.png',
          name: '成龙'
        },
        {
          id: 'chenbaixiang.png',
          name: '陈百祥'
        },
        {
          id: 'zhanzhao.png',
          name: '展昭'
        },
        {
          id: 'chengjisihan.png',
          name: '成吉思汗'
        },
        {
          id: 'hubilie.png',
          name: '忽必烈'
        },
        {
          id: 'yuefei.png',
          name: '岳飞'
        },
        {
          id: 'liuyan.png',
          name: '柳岩'
        },
        {
          id: 'liyifeng.png',
          name: '李易峰'
        },
        {
          id: 'wujing.png',
          name: '吴京'
        },
        {
          id: 'zhangxueyou.png',
          name: '张学友'
        },
        {
          id: 'zhenzidan.png',
          name: '甄子丹'
        },
        {
          id: 'libingbing.png',
          name: '李冰冰'
        },
        {
          id: 'peiyongjun.png',
          name: '裴勇俊'
        },
        {
          id: 'zhoujielun.png',
          name: '周杰伦'
        },
        {
          id: 'zhourunfa.png',
          name: '周润发'
        },
        {
          id: 'guanyu.png',
          name: '关羽'
        },
        {
          id: 'zhugeliang.png',
          name: '诸葛亮'
        },
        {
          id: 'caocao.png',
          name: '曹操'
        },
        {
          id: 'zhouxingchi.png',
          name: '周星驰'
        },
        {
          id: 'liuchuanfeng.png',
          name: '流川枫'
        },
        {
          id: 'chimugangxian.png',
          name: '赤木刚宪'
        },
        {
          id: 'huangfeihong.png',
          name: '黄飞鸿'
        },
        {
          id: 'cangjingkong.png',
          name: '苍井空'
        },
        {
          id: 'mihiro.png',
          name: 'Mihiro'
        },
        {
          id: 'daqiaoweijiu.png',
          name: '大桥未久'
        }, {
          id: 'xiyexiang.png',
          name: '西野翔'
        }, {
          id: 'buzhihuowu.png',
          name: '不知火舞'
        },
        {
          id: 'yingmu.png',
          name: '樱木花道'
        },
        {
          id: 'sunwukong.png',
          name: '孙悟空'
        }, {
          id: 'tangseng.png',
          name: '唐僧'
        }, {
          id: 'shaseng.png',
          name: '沙僧'
        },
        {
          id: 'zhubajie.png',
          name: '猪八戒'
        },
        {
          id: 'torrent.png',
          name: '种子'
        }, {
          id: 'military.png',
          name: '军事'
        },
        {
          id: 'ye.png',
          name: '绿叶'
        },
        {
          id: 'goal.png',
          name: '目标'
        },
        {
          id: 'pen.png',
          name: '笔'
        },
        {
          id: 'taiyang.png',
          name: '太阳'
        },
        {
          id: 'disk.png',
          name: '硬盘'
        },
        {
          id: 'taiyang2.png',
          name: '太阳2'
        }, {
          id: 'lanqiu.png',
          name: '篮球'
        }, {
          id: 'zuqiu.png',
          name: '足球'
        }, {
          id: 'ganlanqiu.png',
          name: '橄榄球'
        },
        {
          id: 'taiyang3.png',
          name: '太阳3'
        },
        {
          id: 'xuehua.png',
          name: '雪花'
        },
        {
          id: 'document.png',
          name: '文件盒'
        },
        {
          id: 'db.png',
          name: '数据1'
        },
        {
          id: 'db2.png',
          name: '数据2'
        },
        {
          id: 'pdf.png',
          name: 'pdf'
        },
        {
          id: 'sugar.png',
          name: '糖果'
        },
        {
          id: 'laptop.png',
          name: '笔记本电脑'
        },
        {
          id: 'chuxuguan.png',
          name: '储蓄罐'
        }, {
          id: 'cpp.png',
          name: 'C++'
        },
        {
          id: 'toolbox.png',
          name: '工具箱'
        },
        {
          id: 'programming.png',
          name: '程序'
        },
        {
          id: 'accbook.png',
          name: '账本'
        },
        {
          id: 'man.png',
          name: '男人'
        },
        {
          id: 'girlfriend.png',
          name: '女孩'
        },
        {
          id: 'movie.png',
          name: '放映机'
        },
        {
          id: 'party.png',
          name: 'Party'
        },
        {
          id: 'threestar.png',
          name: '三星'
        },
        {
          id: 'bomb.png',
          name: '爆炸'
        },
        {
          id: 'threebody.png',
          name: '三体'
        },
        {
          id: 'graphtheory.png',
          name: '图论'
        },
        {
          id: 'nansheng.png',
          name: '男生'
        },
        {
          id: 'chest.png',
          name: '宝箱'
        },
        {
          id: 'airconditioner.png',
          name: '空调'
        }, {
          id: 'mac.png',
          name: 'Mac'
        }, {
          id: 'pai.png',
          name: '数学'
        }, {
          id: 'bodybuilding.png',
          name: '健身'
        }, {
          id: 'list.png',
          name: '列表'
        }, {
          id: 'stats.png',
          name: '统计'
        }, {
          id: 'shell.png',
          name: '贝壳'
        },
        {
          id: 'virus.png',
          name: '病毒'
        }, {
          id: 'linuxshell.png',
          name: 'shell'
        }, {
          id: 'cash.png',
          name: '钱'
        }, {
          id: 'question.png',
          name: '问题'
        }, {
          id: 'python.png',
          name: 'Python'
        }, {
          id: 'php.png',
          name: 'PHP'
        }, {
          id: 'expert.png',
          name: '专家'
        }, {
          id: 'weekend.png',
          name: '周末'
        }, {
          id: 'ubuntu.png',
          name: 'Ubuntu'
        }, {
          id: 'centos.png',
          name: 'Centos'
        }, {
          id: 'github.png',
          name: 'github'
        }, {
          id: 'mysql.png',
          name: 'Mysql'
        }, {
          id: 'docker.png',
          name: 'Docker'
        }, {
          id: 'tensorflow.png',
          name: 'Tensorflow'
        }, {
          id: 'mongo.png',
          name: 'Mongo'
        }, {
          id: 'redis.png',
          name: 'Redis'
        }
      ],
    }
  }
}
