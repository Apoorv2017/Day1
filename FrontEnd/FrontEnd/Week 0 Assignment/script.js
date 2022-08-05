
 
function index(){
    document.entryform.sumbit.addEventListener('click', add_user, false);
}

function add_user(e){
    e.preventDefault();
    console.log("hey");
    let users=JSON.parse(localStorage.getItem('user'))
    if (users==null){
        users=[];
    }
    
    users.push({name:document.entryform.name.value,
                phone:document.entryform.phone.value});

    localStorage.setItem('user', JSON.stringify(users));
    location.href="view.html";
    
}

function showdetails(){
    let users=JSON.parse(window.localStorage.getItem('user'))
    if (users==null){
        users=[];
    }
    let tag="";
    users.forEach(value=>{tag=tag+`<p>Name: ${value.name}, Phone: ${value.phone}</p>`;});
    document.getElementById("showdetails").innerHTML=tag;

}