# 红家后台2023秋季二面题

> Powered by `GoLang 1.19`

## 项目工程性介绍

```
WhoAmI
├── router         路由
│   └── router.go    启动gin服务器，并链接service函数配置路由
├── service        业务函数，主要需要完善该部分的代码
│   ├── difficult    任务四
│   ├── easy         任务二
│   ├── middle       任务三
│   └── picdata    公共数据
├── static         前端
│   ├── assets       素材库
│   ├── data         题目库
│   ├── task1.html   任务一考核前端
│   ├── ...
├── util           公共工具库
│   └── response.go  统一后台返回格式工具
│   └── status.go    统一后台返回错误码
├── go.mod
├── go.sum
├── main.go        入口函数
└── README.md      开发文档
```

## 任务介绍

后台面试小作业共有4个小任务。千万不要被这个数字吓到了，其中只有任务一为必做任务，在此基础上其余任务任挑一道，即你只需完成`任务一` + `任务二 / 任务三 / 任务四 三选一`两个任务。

为减少从零开始的迷茫，提供的项目**已完成大部分基础框架性工作**，任务介绍时会有尽可能详细的任务指引。

当然，你可以对该工程其余部分做任何改动，只不过请让他可以跑起来。

### 【必做】任务一：配置环境，运行项目

#### 任务指引

- 项目基于GoLang开发，需配置GoLang开发环境。
- 为方便开发，建议下载`VSCode`或`GoLand`等编辑器/编译器编辑代码并运行项目
- 为方便测试接口，可下载`ApiFox`或`postman`进行接口测试
- 运行项目

```bash
go run main.go
```

#### 达标要求：

- [ ] 可正常显示[static/test.html](localhost:8080/static/test.html)

![](https://cdn.nlark.com/yuque/0/2022/png/2628060/1663479615534-4d72f3c1-8888-4f22-9295-2ec1206549cc.png)


### 【简单】任务二：提供一个完整的图片列表

补充`/service/easy/query.go`文件中`GetPicDatas`函数，使其返回结果为data文件夹图片及其对应名字

#### 达标要求：

- [ ] `/static/task2.html`网页中正常显示题目并可以正常交互


### 【中等】任务三：提供一道题和答案

补充`/service/middle/query.go`文件中`GetQuestionWithA`函数，使其返回结果为带答案的题目问题

#### 达标要求：

- [ ] `/static/task3.html`网页中正常显示题目并可以正常交互

### 【挑战】任务四：提供一道题，但不提供答案，由后台判断提交的答案是否正确

补充`/service/difficult/query.go`文件中`GetQuestion`与`Verify`函数，使前者返回结果为不带答案的题目，后者验证输入的答案是否正确

#### 达标要求：

- [ ] `/static/task4.html`网页中正常显示题目并可以正常交互

## API接口文档

## 任务一

### GET 测试接口

GET /hello

> 返回示例

> 成功

```json
{
  "code": 200,
  "msg": "成功",
  "data": {
    "text": "hello, world",
    "time": "2022-09-18T14:09:26+0800"
  }
}
```

#### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

#### 返回数据结构

状态码 **200**

| 名称    | 类型    | 必选 | 约束 | 中文名 | 说明 |
| ------- | ------- | ---- | ---- | ------ | ---- |
| » code  | integer | true | none |        | none |
| » msg   | string  | true | none |        | none |
| » data  | object  | true | none |        | none |
| »» text | string  | true | none |        | none |
| »» time | string  | true | none |        | none |

## 任务二

### GET 获取完整图片列表

GET /easy/getList

> 返回示例

> 成功

```json
{
  "code": 200,
  "msg": "成功",
  "data": {
    "picdatas": [
      {
        "filename": "0001.jpeg",
        "name": "迪迦奥特曼"
      },
      {
        "filename": "0002.png",
        "name": "名侦探柯南"
      },
      {
        "filename": "0003.jpg",
        "name": "复仇者联盟"
      },
      {
        "filename": "0004.jpg",
        "name": "哆啦a梦"
      },
      {
        "filename": "0005.jpg",
        "name": "斗罗大陆"
      },
      {
        "filename": "0006.jpg",
        "name": "秦时明月"
      },
      {
        "filename": "0007.jpeg",
        "name": "间谍过家家"
      },
      {
        "filename": "0008.jpg",
        "name": "罗小黑战记"
      },
      {
        "filename": "0009.jpg",
        "name": "快把我哥带走"
      },
      {
        "filename": "0010.jpg",
        "name": "刺客伍六七"
      },
      {
        "filename": "0011.jpg",
        "name": "画江湖之不良人"
      },
      {
        "filename": "0012.jpg",
        "name": "蜡笔小新"
      },
      {
        "filename": "0013.jpg",
        "name": "火影忍者"
      },
      {
        "filename": "0014.jpg",
        "name": "天线宝宝"
      },
      {
        "filename": "0015.jpg",
        "name": "刺客信条"
      }
    ]
  }
}
```

#### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

#### 返回数据结构

状态码 **200**

| 名称         | 类型     | 必选 | 约束 | 中文名 | 说明 |
| ------------ | -------- | ---- | ---- | ------ | ---- |
| » code       | integer  | true | none |        | none |
| » msg        | string   | true | none |        | none |
| » data       | object   | true | none |        | none |
| »» picdatas  | [object] | true | none |        | none |
| »»» filename | string   | true | none |        | none |
| »»» name     | string   | true | none |        | none |

## 任务三

### GET 获取带答案的问题

GET /middle/getQuestion

> 返回示例

> 成功

```json
{
  "code": 200,
  "msg": "成功",
  "data": {
    "filename": "0014.jpg",
    "question": [
      "天线宝宝",
      "复仇者联盟",
      "哆啦a梦",
      "火影忍者"
    ],
    "answer": 0
  }
}
```

#### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

#### 返回数据结构

状态码 **200**

| 名称        | 类型     | 必选 | 约束 | 中文名 | 说明 |
| ----------- | -------- | ---- | ---- | ------ | ---- |
| » code      | integer  | true | none |        | none |
| » msg       | string   | true | none |        | none |
| » data      | object   | true | none |        | none |
| »» filename | string   | true | none |        | none |
| »» question | [string] | true | none |        | none |
| »» answer   | integer  | true | none |        | none |

## 任务四

### GET 获取不带答案的问题

GET /difficult/getQuestion

> 返回示例

> 成功

```json
{
  "code": 200,
  "msg": "成功",
  "data": {
    "id": "Q00001",
    "filename": "0014.jpg",
    "question": [
      "天线宝宝",
      "复仇者联盟",
      "哆啦a梦",
      "火影忍者"
    ]
  }
}
```

#### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

#### 返回数据结构

状态码 **200**

| 名称        | 类型     | 必选 | 约束 | 中文名 | 说明 |
| ----------- | -------- | ---- | ---- | ------ | ---- |
| » code      | integer  | true | none |        | none |
| » msg       | string   | true | none |        | none |
| » data      | object   | true | none |        | none |
| »» id       | string   | true | none |        | none |
| »» filename | string   | true | none |        | none |
| »» question | [string] | true | none |        | none |

### POST 验证选项是否正确

POST /difficult/verify

> Body 请求参数

```yaml
id: q0101
filename: 0001.jpg
chosenId: "0"
chosenText: A
chosenName: 天线宝宝

```

#### 请求参数

| 名称         | 位置 | 类型    | 必选 | 说明     |
| ------------ | ---- | ------- | ---- | -------- |
| body         | body | object  | 否   | none     |
| » id         | body | string  | 否   | 问题ID   |
| » filename   | body | string  | 否   | 文件名   |
| » chosenId   | body | integer | 否   | 选项ID   |
| » chosenText | body | string  | 否   | 选项字母 |
| » chosenName | body | string  | 否   | 选项名   |

> 返回示例

> 成功

```json
{
  "code": 200,
  "msg": "成功",
  "data": {
    "status": true
  }
}
```

```json
{
  "code": 200,
  "msg": "成功",
  "data": {
    "status": false
  }
}
```

#### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

#### 返回数据结构

状态码 **200**

| 名称      | 类型    | 必选 | 约束 | 中文名 | 说明 |
| --------- | ------- | ---- | ---- | ------ | ---- |
| » code    | integer | true | none |        | none |
| » msg     | string  | true | none |        | none |
| » data    | object  | true | none |        | none |
| »» status | boolean | true | none |        | none |
