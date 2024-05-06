//to dispaly all the added students
window.onload = function(){
    fetch("/student")
    .then(response => response.text())
    .then(data => showStudents(data));
}

//
function addRow(student){
    var table = document.getElementById("myTable");
    var row = table.insertRow(table.length);
    // Insert new cells (<td> elements) at the 1st and 2nd position of the"new" <tr> element:
    var td=[]
    for(i=0; i<table.rows[0].cells.length; i++){
    td[i] = row.insertCell(i);
    }
    // Add student detail to the new cells:
    td[0].innerHTML = student.stdid;
    td[1].innerHTML = student.fname;
    td[2].innerHTML = student.lname;
    td[3].innerHTML = student.email;
    td[4].innerHTML = '<input type="button" onclick="deleteStudent(this)"value="delete" id="button-1">';
    td[5].innerHTML = '<input type="button" onclick="updateStudent(this)"value="edit" id="button-2">';
}


function showStudents(data){
        const students = JSON.parse(data)
        students.forEach(stud => {
            var table = document.getElementById("myTable");
        addRow(stud)
    })

}

function showStudent(data) {
    const student = JSON.parse(data)
    addRow(student)
    }
    
// grtfromdata
function getFormData(){
    var data = {
        stdid : parseInt(document.getElementById("sid").value),
        fname : document.getElementById("fname").value,
        lname : document.getElementById("lname").value,
        email : document.getElementById("email").value
        }
        return data
}

//adding student
function addStudent(){  
    var data = getFormData()
    fetch('/student', {
    method: "POST",
    body: JSON.stringify(data),
    headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(response1 =>{
        var sid = data.stdid;
        if(response1.ok){
            fetch('/student/'+sid)
            .then(response2 => response2.text())
            .then(data => showStudent(data))
    }else{
        throw new Error(response1.status)
    }
}).catch(e =>{
    if (e.message ==303){
        alert("user not logged in.")
        window.open("index.html","_self")
    }else if(e.message == 500){
        alert("server error!")
    }
});

resetform(); 
    var sid = data.stdid
    if (isNaN(sid)){
        alert("Enter valid student ID")
        return
    }else if (data.email == ""){
        alert("Email cannot be empty")
        return
    }else if (data.fname === ""){
        alert("first name cannot be empty")
        return
    }
    console.log(data)

    
}

//reset form
function resetform(){
    document.getElementById("sid").value = "";
    document.getElementById("fname").value = "";
    document.getElementById("lname").value = "";
    document.getElementById("email").value = "";
}

//delete button
function deleteStudent(r){
    // this(input) -> td -> tr
    if (confirm('Are you sure you want to DELETE this?')){
    selectedRow = r.parentElement.parentElement;
    sid = selectedRow.cells[0].innerHTML;
    fetch('/student/'+sid, {
    method: "DELETE",
    headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(res => {
        if(res.ok){
            alert("student deleted")
            var rowIndex = selectedRow.rowIndex
            if(rowIndex){
                document.getElementById("myTable").deleteRow(rowIndex)
            }
            selectedRow = null
        }
    })
   }
}

function update(sid){
    //extract new data from the form
    var newData = getFormData()
    fetch("/student/"+sid, {
        method: "PUT",
        //json stringify convert js object into go object
        body: JSON.stringify(newData),
        // headers: {"Content-type":"application/json; charset=UTF-8"}
    }).then(res => {
        if (res.ok){
            //fill in selected row with updated value
            selectedRow.cells[0].innerHTML = newData.stdid;
            selectedRow.cells[1].innerHTML = newData.fname;
            selectedRow.cells[2].innerHTML = newData.lname;
            selectedRow.cells[3].innerHTML = newData.email;

            //set to previous value
            var button = document.getElementById("button-add");
            button.innerHTML = "Add";
            button.setAttribute("onclick", "addStudent()");
            selectedRow = null;

            resetform();

        }else{
            alert("server: Update request error.")
        }
    })
}
//update button
var selectedRow = null
function updateStudent(r){

    //r.parentElement is td or data is stored in td
    selectedRow = r.parentElement.parentElement;

    //filling the form as soon as we click on edit button and update
    document.getElementById("sid").value = selectedRow.cells[0].innerHTML;
    document.getElementById("fname").value = selectedRow.cells[1].innerHTML;
    document.getElementById("lname").value = selectedRow.cells[2].innerHTML;
    document.getElementById("email").value = selectedRow.cells[3].innerHTML;  
    
    //
    var btn = document.getElementById("button-add")
    if(btn){
        btn.innerHTML = "update"
        sid = selectedRow.cells[0].innerHTML
        btn.setAttribute("onclick","update(sid)")
    }
}
