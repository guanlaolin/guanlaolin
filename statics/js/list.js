onload = function(){
  GetInfo();
}

function GetInfo(){
  if(window.ActiveXObject){ //如果是IE浏览器
    var xmlHttpRequest = new ActiveXObject("Microsoft.XMLHTTP");
  }else if(window.XMLHttpRequest){ //非IE浏览器
    var xmlHttpRequest = new XMLHttpRequest();
  }

  //回调函数
  xmlHttpRequest.onreadystatechange = function(){
    if (xmlHttpRequest.readyState==4 && xmlHttpRequest.status == 503){
      alert("获取文件信息失败,请刷新浏览器或联系管理员!");
      return;
    }

    if (xmlHttpRequest.readyState==4 && xmlHttpRequest.status == 200){

    var info = document.getElementById("info");

    var obj = xmlHttpRequest.responseText;
    var jsonObj = eval('('+obj+')');
    //防止jsonObj为空
    if (jsonObj==null){
      info.innerHTML="<h1>无文件</h1>";
    }
    //alert(jsonObj);
    for (var i=0;i<jsonObj.length;i++){
      var url = jsonObj[i].url;
      var name = jsonObj[i].name;
      var time = jsonObj[i].time;
      var tr = document.createElement("tr");
      tr.innerHTML = "<td>"+name+"</td><td>"+time+"</td><td><a href=/download?id="+url+">下载</a></td>";
      info.appendChild(tr);
    }

  }
  };

  xmlHttpRequest.open("GET","/list");
  xmlHttpRequest.send();

}

function showFileDiag(){
  document.getElementById("file").style.visibility = "visible";
}

function hiddenFileDiag(){
  document.getElementById("file").style.visibility = "hidden";
}
