<html>
<head>
<title></title>
</head>
<body onload="redirect()">
  <div>
    <span style='color:red'> {{.msg}} </span>, <span id='sec'>3</span>秒后<a href="javascript:history.back()">返回</a>
    </div>
    <script type="text/javascript">
    function redirect(){
      var t=2;
      setInterval(function(){
      document.getElementById('sec').innerHTML=t;
            t--;
        },1000);
        setTimeout(function  () {
         	history.back()
        }, 3000);
    }

    </script>
</body>
</html>