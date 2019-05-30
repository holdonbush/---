//cal.js
var wxcharts = require("../../utils/wxcharts.js")
var util = require("../../utils/util.js")
const app = getApp()
var columnChart = null;
var sliderWidth = 96; // 需要设置slider的宽度，用于计算中间位置
Page({
  global:{
    nickName:"111",
    d1:"",
    d2:"",
    d3:"",
    d4:"",
    d5:"",
    d6:"",
    d7:"",
    n0:"",
    n1:"",
    n2:"",
    n3:"",
    n4:"",
    n5:"",
    n6:"",
    total:0,
  },
  data: {
    tabs: ["最近的消费", "今日消费记录", "图表"],
    names: global.nickName,
    activeIndex: 1,
    sliderOffset: 0,
    sliderLeft: 0
  },
  backToMainChart: function () {
    this.setData({
      chartTitle: "消费情况图",
      isMainChartDisplay: false
    });
    columnChart.updateData({
      series: [{
        name: '成交量',
        data: [global.n0, global.n1, global.n2, global.n3, global.n4, global.n5, global.n6]
      }]
    });
  },
  onLoad: function () {
    wx.getUserInfo({
      withCredentials: true,
      lang: '',
      success: function(res) {
        global.nickName = res.userInfo.nickName
      },
      fail: function(res) {},
      complete: function(res) {},
    })
    var data1 = {name : app.globalData.userInfo.nickName}
    var that = this;
    wx.request({
      url: 'http://120.79.240.163:9090/thisday',
      data: data1,
      method:"POST",
      success: function(e) {
        global.total = e.data.Total
        that.setData({
          thisday:e.data.Total,
          yest:e.data.YesterdayT,
          lastWeek:e.data.YesterdayT+e.data.Total+e.data.D3t+e.data.D4t+e.data.D5t+e.data.D6t+e.data.D7t,
          chartTitle:"消费情况图",
          isMainChartDisplay: false
        })
        console.log("get thisday info success")
        console.log(e)
        console.log(e.data.Total)
        var time = util.formatDate(new Date());
        var time1 = util.formatDateP(new Date, -1);
        global.d1 = time
        global.d2 = time1
        global.d3 = util.formatDateP(new Date, -2);
        global.d4 = util.formatDateP(new Date, -3);
        global.d5 = util.formatDateP(new Date, -4);
        global.d6 = util.formatDateP(new Date, -5);
        global.d7 = util.formatDateP(new Date, -6);
        console.log(global.d1, global.d2, global.d3, global.d4, global.d5, global.d6, global.d7)
        global.n0 = e.data.Total
        global.n1 = e.data.YesterdayT
        global.n2 = e.data.D3t
        global.n3 = e.data.D4t
        global.n4 = e.data.D5t
        global.n5 = e.data.D6t
        global.n6 = e.data.D7t
        columnChart = new wxcharts({
          canvasId: 'columnCanvas',
          type: 'column',
          categories: [global.d1,global.d2,global.d3,global.d4,global.d5,global.d6,global.d7],
          series: [{
            name: '成交量1',
            data: [e.data.Total,e.data.YesterdayT,e.data.D3t,e.data.D4t,e.data.D5t,e.data.D6t,e.data.D7t]
          }],
          yAxis: {
            format: function (val) {
              return val + '元';
            }
          },
          width: 320,
          height: 200
        });
      
      }
    })
    wx.getSystemInfo({
      success: function (res) {
        that.setData({
          sliderLeft: (res.windowWidth / that.data.tabs.length - sliderWidth) / 2,
          sliderOffset: res.windowWidth / that.data.tabs.length * that.data.activeIndex
        });
      }
    });
  },
  tabClick: function (e) {
    this.setData({
      sliderOffset: e.currentTarget.offsetLeft,
      activeIndex: e.currentTarget.id
    });
    if(e.currentTarget.id == 2) {
      console.log(1)
      this.setData({
        chartTitle: "消费情况图",
        isMainChartDisplay: false
      });
      columnChart.updateData({
        series: [{
          name: '成交量',
          data: [global.n0, global.n1, global.n2, global.n3, global.n4, global.n5, global.n6]
        }]
      });
    }
  },
  submit:function(e) {
    e.detail.value.name = global.nickName
    console.log(e.detail.value)
    var that = this;
    wx.request({
      url: 'http://120.79.240.163:9090/d1',
      method:"POST",
      data:e.detail.value,
      success:function(res) {
        if(e.detail.value.num == "") {
          global.total = global.total + 0
        }
        else {
          global.total = global.total + parseInt(e.detail.value.num)
        }
        console.log(global.total)
        that.setData({
          thisday:res.data.Total,
          form_info:''
        })
        global.n0 = res.data.Total
        global.n1 = res.data.YesterdayT
        global.n2 = res.data.D3t
        global.n3 = res.data.D4t
        global.n4 = res.data.D5t
        global.n5 = res.data.D6t
        global.n6 = res.data.D7t
      }
    })
  },
});