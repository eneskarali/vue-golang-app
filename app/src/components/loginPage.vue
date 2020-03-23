<template>
<div class="page">
    <div class="login">  
        <h1>Kullanıcı Bilgileri</h1>      
        <ul>
           <li><p>Kullanıcı Adı:</p><input type="text"/></li>
           <li><p>Şifre:</p><input type="password"/></li>
        </ul>          
    </div>
</div>
</template>

<script>
export default {
  name: 'loginPage',
  props: {
    msg: String
  }
}
</script>

<style scoped>
p{
    float: left;   
    font-size: 22px;
    margin:2px;
    color: black;
    font: normal;
}
h1 {
    margin-top:10px;
    margin-bottom: 15px;
    color: black;
}
ul {
    list-style-type: none;
    padding: 0;
    margin: 0;
    border-style: solid;
    border-color: #4caf50;
    border-block-width: 20px;
    border-inline-width: 15px;
    border-radius: 10px;
    background-color: #4caf50;

}
li {
    margin: 15px 10px;
}
input{
    height: 35px;
    width: 98%;
    font-size: 20px;
}
.login{
    margin-top: 70px;
    text-align: center;
    width: 550px;
    display: inline-block;
    
}

.page{
    text-align: center;
}
@media(max-width: 800px){
    input{
        width: 320px;
    }

}


</style>