# 内嵌helm
### 目录
- 一. 规则声明
- 二. chart目录
- 三. chart自定义参数
- 四. template常用语法


## 一. 规则声明
- 所有chart包需经过校验,才能加入置放到 runtime/charts 下
- chart目录在 runtime/charts/index.yaml
- chart命名规则为: {{chart_name}}@{{version}}
- chart自定义参数不能与chart中的values.yaml冲突

## 二. chart目录
```yaml
Chart.yaml          # Yaml文件，用于描述Chart的基本信息，包括名称版本等
LICENSE             # [可选] 协议
README.md           # [可选] 当前Chart的介绍
values.yaml         # Chart的默认配置文件
requirements.yaml   # [可选] 用于存放当前Chart依赖的其它Chart的说明文件
charts/             # [可选]: 该目录中放置当前Chart依赖的其它Chart
templates/          # [可选]: 部署文件模版目录，模版使用的值来自values.yaml和由Tiller提供的值
templates/NOTES.txt # [可选]: 放置Chart的使用指南
```


## 三. chart自定义参数
### 1. Bool按钮
- type 参数类型: radio
- label 标签名
- key 参数键
- required 是否必选
- default 默认值: true/false
- disabed 禁止修改: true (仅当有默认值时生效)

例子:
```json
{
    "type": "radio",
    "label": "是否打开外部访问",
    "key": "is_open_ingress",
    "required": true
}
```

### 2. 文本输入框
- type 参数类型: text
- label 标签名
- key 参数键
- unit 参数值单位
- required 是否必选
- placeholder 提示占用符
- default 默认值
- disabed 禁止修改: true (仅当有默认值时生效)
- max 文本最大长度
- rules 文本规则组
    - pattern 规则: 正则表达式
    - message 提示信息
    
例子:
```json
{
    "type": "text",
    "label": "端口号",
    "key": "port",
    "placeholder": "端口范围为1000-30000",
    "default": "2000",
    "max": 5,
    "rules": [{
      "pattern": "/^[0-9]*$/",
      "message": "端口号为纯数字"
    }]
}
```


### 3. 分组
- type 参数类型:  group
- label 标签名
- values 对象组 (可包含其他参数类型)

例子:
```json
{
    "type": "group",
    "label": "配置",
    "values": [{
        "type": "text",
        "label": "内存",
        "key": "memory",
        "unit": "Mi"
    },{
        "type": "radio",
        "label": "是否保存",
        "key": "is_save",
        "required": true
    }]
}
```

### 4. 选择框
- type 参数类型:  select
- label 标签名 
- key 参数键值 
- required 是否必选 
- disabed 禁止修改: true (仅当有默认选项时生效)
- default 默认选项,选项的value  
- options 选择组 
    - value 选择项 
    - label 选择标签名 

例子:
```json
{
    "type": "select",
    "label": "选择集群类型",
    "key": "cluster_type",
    "default": "one",
    "options": [{
        "label": "单节点",
        "value": "one"
    }, {
        "label": "主从集群",
        "value": "cluster"
    }]
}
```

## 四. template常用语法

> {{ .Values.* }}
> 
> 从value.yaml文件中读取

>{{ .Release.* }}
>
>从运行Release的元数据读取

>{{ .Template.* }}   {{ .Chart.* }}
>
>从Chart.yaml文件中读取


>{{ .Files.* }}
>
>文件数量少情况下,在chart的根目录下有三个文件,在模板文件中使用

>{{ .Capabilities.* }}
>
>文件多情况下使用


>{{quote }}
>
>{{ quote  .Values.favorite.drink }}
>
>是最常用的模板函数，它能把ABC转化为“ABC”。它带一个参数

>{{ template }}  {{ include  }}
>- 1、 先在_helpers.tpl中定义命名模板
>- 2、使用命名模版 
>- 3、渲染后 


>{{ |default }}
> 
>drink: {{ .Values.favorite.drink | default “tea” | quote }}
> 
>如果在values中无法找到favorite.drink，则配置为“tea”。

>{{  |indent }}
> 
>{{ include "mychart_app" . | indent 2 }}
> 
>对左空出空格