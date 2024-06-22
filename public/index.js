
 function removeFromDb(item){
   fetch(`/delete?item=${item}`, {method: "Delete"}).then(res =>{
       if (res.status == 200){
           window.location.pathname = "/"
       }
   })
}

function updateDb(olditem) {
   let input = document.getElementById(olditem)
   let newitem = input.value
   console.log(olditem, newitem);
   fetch(`/update/${olditem}/${newitem}`, {method: "PUT"}).then(res =>{
       if (res.status == 200){
       alert("Database updated")
           window.location.pathname = "/"
       }
   })
}

