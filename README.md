# GezginWebEngineGo
Gezgin Web Engine written in Go Completely From Scratch

# It's and Experimental Browser Engine

![image](https://user-images.githubusercontent.com/16630690/232159321-510a4fcf-b080-4835-b453-dcf45abb2408.png)

Html File Showed Above As Screenshot
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <title>Title</title>
    <script>
        const colors = ["blue", "red", "green", "cyan", "darkblue"];
        function change_div_background_color() {
            let random = 0
            random = Math.floor(Math.random() * colors.length);
            document.getElementById("div1").setAttribute("style", "background-color:"+colors[random])
            random = Math.floor(Math.random() * colors.length);
            document.getElementById("div2").setAttribute("style", "background-color:"+colors[random])
            random = Math.floor(Math.random() * colors.length);
            document.getElementById("div3").setAttribute("style", "background-color:"+colors[random])
            setTimeout(change_div_background_color, 3000);
        }
        setTimeout(change_div_background_color, 3000);
    </script>
    <style>
        .mydiv {background-color: red; position: relative; left: 30px; top: 30px}
    </style>
</head>
<body>
    <div style = "position:relative; left:20px; top:20px; background-color: black">
        <div id="div1" style ="background-color:darkblue">
            Hello World
        </div>
        <div id="div2" style=" background-color:green">
            This is new Hello World
        </div>
        <div id="div3 " class="mydiv">
            This is 3rd div
            <div style=" background-color:red ">
                This is an inside div
            </div>
        </div>
    </div>
</body>
</html>
```

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

### It uses SDL 2.0 for drawing(It will be replaced by OpenGL or Skia)
### It uses v8 js-engine for executing javascript code
