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


# Supported Html Elements for Parsing

a  
button  
canvas  
details  
img  
label  
link  
script  
style  
textarea  

# Supported Html Elements for Drawing
html  
body  
div  

# Supported Css Properties
accent-color  
align-content  
align-items  
align-self  
background-blend-mode  
background-repeat  
background-origin  
background-clip  
background-attachment  
background-color  
border  
color  
height  
min-height  
max-height  
width  
min-width  
max-width  
margin-top  
margin-bottom  
margin-left  
margin-right  
margin  
padding-top  
padding-bottom  
padding-left  
padding-right  
padding  
position  
top  
bottom  
left  
right  
visibility  

### It uses SDL 2.0 for drawing(It will be replaced by OpenGL or Skia)
### It uses v8 js-engine for executing javascript code
