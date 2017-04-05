<html>
<head>
<title>success</title>
</head>
<body onload="redirect()">
  <div>
       <span style='color:green'> {{.msg}}</sapn> , <span id='sec'>1</span>秒后<a href="{{.url}}">跳转</a>
    </div>
    <script type="text/javascript">
    function redirect(){
      var t=1;
      setInterval(function(){
      document.getElementById('sec').innerHTML=t;
            t--;
        },1000);
        setTimeout(function  () {
          location.href="{{.url}}"
        }, 500);
    }
    </script>
</body>
</html>