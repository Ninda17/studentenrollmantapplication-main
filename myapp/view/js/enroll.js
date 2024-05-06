window.onload = function(){
    //request to get all students
    fetch("/student")
    .then(response => response.text())
    .then(data => getStudents(data));

    //get all course
    fetch("/courses")
    .then(response => response.text())
    .then(data => getCourses(data));

    //get all enrollment
    fetch('/enrolls')
    .then(response => response.text())
    .then(data => getAllEnroll(data));

}

function getStudents(data){
    const students = [] // to store all student ID
    const allStudents = JSON.parse(data)
    // add student id to students list
    allStudents.forEach(stud => {
        students.push(stud.stdid)
    });
    // get the sid input field
    var select = document.getElementById("sid")
// iterate over the students and create a new option for each ID
    for (var i = 0; i < students.length; i++) {
        var sid = students[i];
        var option = document.createElement("option")
        option.textContent= sid;
        option.value = sid;
        select.appendChild(option);
    }
}

//drop down
function getCourses(data){
    const courses = []
    const allCourses = JSON.parse(data)
    allCourses.forEach(course =>{
        courses.push(course.courseid)
    });

    var option = "";
    for(var i = 0; i < courses.length; i++){
        option += '<option value ="'+ courses[i]+'">'+courses[i]+"</option>"
    }
     document.getElementById("cid").innerHTML= option;
}



//add enrollment
function addEnroll() {
    var _data = {
        stdid : parseInt(document.getElementById("sid").value),
        cid : document.getElementById("cid").value,
    }

    var sid = _data.stdid;
    var cid = _data.cid;

    if (isNaN(sid) || cid == ""){
        alert("Slect valid data")
        return
    }
    fetch("/enroll",{
        method:"POST",
        body: JSON.stringify(_data),
        headers:{"content-type":"application/json; charset=UTF-8"}
    }).then(response => {
        if (response.ok) {
            fetch('/enroll/'+sid+'/'+cid)
            .then(response => response.text())
            .then(data => getEnrolled(data))
        } else {
            throw new Error(response.statusText)
        }
    }).catch(e => {
        if (e == "Error: Forbidden") {
            alert(e+". Duplicate entry!")
        }
    });
        resetFields();

}

//show enrollment if the post method is successful
function getEnrolled(data) {
    const enrolled = JSON.parse(data)
    showTable(enrolled)
}

//helper function to add new row to dispaly a data od single student
function showTable(enrolled) {
    // Find a <table> element with id="myTable":
    var table = document.getElementById("myTable");
    // Create an empty <tr> element and add it to the last position of the table:
    var row = table.insertRow(table.length);
    // Insert new cells (<td> elements) at the 1st and 2nd position of the "new" <tr> element:
    var td=[]
    for(i=0; i<table.rows[0].cells.length; i++){
        td[i] = row.insertCell(i);
    }
    td[0].innerHTML = enrolled.stdid;
    td[1].innerHTML = enrolled.cid;
    td[2].innerHTML = enrolled.date.split("T")[0]; // show only date, ignore time
    td[3].innerHTML = '<input type="button" onclick="deleteEnroll(this)" value="Delete" id="button-1">';
    
}

//reset
function resetFields() {
    document.getElementById("sid").value = "";
    document.getElementById("cid").value = "";
}

//get all enrollment
function getAllEnroll(data) {
    const allenroll = JSON.parse(data)
    allenroll.forEach(enroll => {
        showTable(enroll)
    });
}


//delete enrollment
const deleteEnroll = async(r) => {
    if (confirm('Are you sure you want to DELETE this?')){
        selectedRow = r.parentElement.parentElement;
        sid = selectedRow.cells[0].innerHTML;
        cid = selectedRow.cells[1].innerHTML;
        fetch('/enroll/'+sid+"/"+cid, {
            method: "DELETE",
            headers: {"Content-type": "application/json; charset=UTF-8"}
    }). then( response => {
        if (response.ok) {
            var rowIndex = selectedRow.rowIndex; // index starts from 0
            if (rowIndex>0) { //th is row 0
                document.getElementById("myTable").deleteRow(rowIndex);
            }
        }
    });
}
}
    
