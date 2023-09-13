# GezginWebEngineGo
Gezgin Web Engine written in Go Completely From Scratch

# It's an Experimental Browser Engine

![Screenshot from 2023-09-13 19-33-11](https://github.com/MrSametBurgazoglu/GezginWebEngineGo/assets/16630690/965a7872-dd3c-4243-bd81-b2155c04d09c)

https://getbootstrap.com/docs/5.0/examples/cover/ Web Adress Showed Above As Screenshot

### WebSupport
Now Gezgine Can Load Web Pages From Web

### Run
```
git clone https://github.com/MrSametBurgazoglu/GezginWebEngineGo.git
cd GezginWebEngineGo
go get .
go run main/main.go
```


# Supported Html Elements

| Html Elements    | Parsing | Drawing |
|------------------|---------|---------|
| ```<a>```        | &#9745; |         |
| ```<body>```     | &#9745; | &#9745; |
| ```<button>```   | &#9745; |         |
| ```<canvas>```   | &#9745; |         |
| ```<details>```  | &#9745; |         |
| ```<div>```      | &#9745; | &#9745; |
| ```<html>```     | &#9745; | &#9745; |
| ```<img>```      | &#9745; |         |
| ```<label>```    | &#9745; |         |
| ```<link>```     | &#9745; |         |
| ```<script>```   | &#9745; |         |
| ```<style>```    | &#9745; |         |
| ```<textarea>``` | &#9745; |         |

# Supported Css Properties


| Css Properties               | Parsing | In-Use  |
|------------------------------|---------|---------|
| ```accent-color```           | &#9745; |
| ```align-content```          | &#9745; |
| ```align-items ```           | &#9745; |
| ```align-self```             | &#9745; |
| ```background-blend-mode ``` | &#9745; |
| ```background-repeat```      | &#9745; |
| ```background-origin```      | &#9745; |
| ```background-clip```        | &#9745; |
| ```background-attachment```  | &#9745; |
| ```background-color```       | &#9745; | &#9745; |
| ```border```                 | &#9745; |
| ```color```                  | &#9745; | &#9745; |
| ```height```                 | &#9745; |
| ```min-height```             | &#9745; |
| ```max-height```             | &#9745; |
| ```width```                  | &#9745; |
| ```min-width```              | &#9745; |
| ```max-width```              | &#9745; |
| ```margin-top```             | &#9745; |
| ```margin-bottom```          | &#9745; |
| ```margin-left```            | &#9745; |
| ```margin-right```           | &#9745; |
| ```margin```                 | &#9745; |
| ```padding-top```            | &#9745; |
| ```padding-bottom```         | &#9745; |
| ```padding-left```           | &#9745; |
| ```padding-right```          | &#9745; |
| ```padding```                | &#9745; |
| ```position```               | &#9745; | &#9745; |
| ```top```                    | &#9745; |
| ```bottom```                 | &#9745; |
| ```left```                   | &#9745; |
| ```right```                  | &#9745; |
| ```visibility```             | &#9745; |

### It uses default go draw library (It will be replaced by OpenGL or Skia)
### It uses Gtk4 for simple window management (For simple web-browser experience)
### It uses v8 js-engine for executing javascript code
