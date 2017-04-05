<html>
<head>
<title>exec</title>
</head>
<body onload="redirect()">

  <div>
      {{.msg}} , <span id='sec'>8</span>秒后返回,<a href="{{.url}}">马上跳转</a>
      <Br>
    </div> 
     {{.str}}
    <script type="text/javascript">
    function redirect(){
      var t=7;
      setInterval(function(){
      document.getElementById('sec').innerHTML=t;
            t--;
        },1000);
        setTimeout(function  () {
          location.href="{{.url}}"
        },8000);
    }
    </script>
</body>
</html>