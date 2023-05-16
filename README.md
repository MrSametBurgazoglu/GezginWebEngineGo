# GezginWebEngineGo
Gezgin Web Engine written in Go Completely From Scratch

# It's and Experimental Browser Engine

![Screenshot from 2023-05-16 22:22:19](https://github.com/MrSametBurgazoglu/GezginWebEngineGo/assets/16630690/54615d1e-d1e2-403d-99e5-da99c9538410)

Html File Showed Above As Screenshot
```html
<!DOCTYPE html>
<html lang="en">
<body style="background-color:rgb(33, 37, 41); color: white">
<div style="position:relative;left:300px">
<h1>What is browser engine and how does it work?</h1>
</div>

<div style="position: absolute; left:20px; top:50px">
<img alt="What is browser engine and how does it work?" width="413" height="373" src="browser-diagram.png">
</div>

<div style="position: absolute; left:470px; top:50px; width: 800px">
<h2>Name and scope</h2>
<p>A browser engine is not a stand-alone computer program but a critical piece of a more extensive program, such as a web browser, from which the term is derived. The word "engine" is an analogy to the engine of a car.</p>

<p>Besides "browser engine", two other terms are in everyday use regarding related concepts: "layout engine" and "rendering engine".In theory, layout and rendering (or "painting") could be handled by different engines. In practice, however, they are tightly coupled and rarely considered separately.<p>

<p>In addition to layout and rendering, a browser engine enforces the security policy between documents, handles navigation through hyperlinks and data submitted through forms, and implements the Document Object Model (DOM) data structure exposed to page scripts</p>

<p>Executing JavaScript (JS) code is a separate matter, however, as every significant web browser uses a dedicated engine for this. The JS language was initially created for use in browsers, but it is now used elsewhere, too, so the implementation of JS engines is decoupled from browser engines. The two engines work in concert via the shared DOM data structure in a web browser.</p>

<h2>Layout and rendering</h2>

<p>The layout of a web page is typically specified by Cascading Style Sheets (CSS). Each style sheet is a series of rules which the browser engine interprets. For example, some rules specify typography details, such as font, color, and text size. The engine combines all relevant CSS rules to calculate precise graphical coordinates for the visual representation it will paint on the screen.</p>

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
