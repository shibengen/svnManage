<html>
<head>
<title>项目列表</title>
<style type="text/css">
  .table th{
    background-color: #ccc
  }
  .table tr{
    background-color: #eee;
    line-height: 30px
  }
  a:hover{
    color: #ff9900
  }  

  a:link{
    color: blue
  }
  a:active{
    color: blue
  }
</style>
</head>
<body>
  <header>
    <h1 class="logo">当前账号:{{.Username}}   <span id='date'></span> <a href="/logout">退出</a></h1>
    <div class="description">
      <!-- <a href="/add">添加</a> -->
    <Table border='1' class='table'>
    <th>项目路径</th><th>备注</th><th>操作</th>
    {{range $k,$v := .list}}
    <Tr><td width="300">
     {{$v.Path}}</td><td width="300">  {{$v.Info}} </td><td><a  href="/exec/?id={{$v.Id}}">更新SVN</a> &nbsp;&nbsp;<a href="/edit/?id={{$v.Id}}">编辑路径</a>&nbsp;&nbsp;<a target='_blank' href="/log/{{$v.Id}}.log">日志记录</a>
   </td></tr>
    {{end}} 
    </table>
    </div>
  </header>
  <footer>
    <div class="author">
     
    </div>
  </footer>
  <div class="backdrop"></div>
</body>
</html>
<script>

</script>