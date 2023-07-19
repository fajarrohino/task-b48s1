// let name = "fajar rohino";
// console.log(`nama saya ${name}`);

// validate alert
// function validForm() {
//   //   event.preventDefault();
//   let projectName = document.getElementById("input-project-name").value;
//   let startDate = document.getElementById("input-start-date").value;
//   let endDate = document.getElementById("input-end-date").value;
//   let description = document.getElementById("input-description").value;
//   let multichecked = document.querySelectorAll(".input-multichecked:checked");
//   let file = document.getElementById("input-file").value;

//   if (projectName === "") {
//     return alert("Nama Project harus diisi!");
//   } else if (startDate === "") {
//     return alert("Waktu Pertama harus diisi!");
//   } else if (endDate === "") {
//     return alert("Waktu berakhir harus diisi!");
//   } else if (description === "") {
//     return alert("Deskripsi harus diisi!");
//   } else if (multichecked.length === 0) {
//     return alert("Harus memilih technologies!");
//   } else if (file === "") {
//     return alert("Upload image harus diisi!");
//   }
// }

let dataProject = [];
function addProject(event) {
  event.preventDefault();
  let projectName = document.getElementById("input-project-name").value;
  let startDate = document.getElementById("input-start-date").value;
  let endDate = document.getElementById("input-end-date").value;
  let description = document.getElementById("input-description").value;
  let nodeJs = document.getElementById("input-nodejs").checked;
  let reactJs = document.getElementById("input-reactjs").checked;
  let nextJs = document.getElementById("input-nextjs").checked;
  let typoScript = document.getElementById("input-typoscript").checked;
  let image = document.getElementById("input-blog-image").files;

  // waktu
  let distanceProject = "";
  let startTgl = new Date(startDate);
  let endTgl = new Date(endDate);

  let distanceTime = endTgl.getTime() - startTgl.getTime(); //default nailai detik
  let distanceDay = distanceTime / (1000 * 3600 * 24); //rentan waktu 1000 milidetik dikali satu detik perjam dan dikali 24 jam
  let distanceWeek = Math.round(distanceDay / 7);
  let distanceMonth = Math.round(distanceWeek / 4);
  let distanceYear = Math.round(distanceWeek / 12);

  if ((distanceDay >= 1) & (distanceDay < 7)) {
    distanceProject = distanceDay + " Days";
  } else if ((distanceWeek >= 1) & (distanceWeek < 4)) {
    distanceProject = distanceWeek + " Week";
  } else if ((distanceMonth >= 1) & (distanceMonth < 12)) {
    distanceProject = distanceMonth + " Month";
  } else if (distanceYear >= 1) {
    distanceProject = distanceYear + " Year";
  } else {
    distanceProject = "Project Expired";
  }

  let file = URL.createObjectURL(image[0]);

  let project = {
    projectName,
    distanceProject,
    description,
    nodeJs,
    reactJs,
    nextJs,
    typoScript,
    file,
    author: "eannoy",
  };

  dataProject.push(project);
  renderProject();
  console.log(dataProject);
}
function renderProject() {
  document.getElementById("contents-project").innerHTML = "";
  for (let index = 0; index < dataProject.length; index++) {
    document.getElementById("contents-project").innerHTML += `
        <div class="list1">
            <div class="project-img">
              <img src="${dataProject[index].file}" alt="" />
            </div>
            <div class="project-content">
              <h4>
                <a href="project-detail.html">${dataProject[index].projectName}</a>
              </h4>
              <div class="detail-project">
                <span>durasi : ${dataProject[index].distanceProject}</span>
                <br /><br />
                <p>${dataProject[index].description}</p>
              </div>
              <div class="framework">
                ${dataProject[index].nodeJs ? "<i class='fa-brands fa-node fa-2x' id='nodejs'></i>" : ""}
                ${dataProject[index].reactJs ? "<i class='fa-brands fa-react fa-2x' id='reactjs'></i>" : ""}
                ${dataProject[index].nextJs ? "<i class='fa-solid fa-n fa-2x' id='nextjs'></i>" : ""}
                ${dataProject[index].typoScript ? "<i class='fa-solid fa-t fa-2x' id='typoscript'></i>" : ""}
              </div>
              <div class="btn-group">
                <button>Edit</button>
                <button>Delete</button>
              </div>
            </div>
          </div>`;
  }
}
