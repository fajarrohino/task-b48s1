// const xhr = new XMLHttpRequest(); //variabel untuk apa ini? -> object can be used to exchange data with a web server

// // CRUD
// xhr.open("GET", "https://your-url", true); // asynchronous
// // param 1 : is the method
// // param 2 : place of data by url
// // param 3 : true -> asynchronous, false -> synchronous

// xhr.onload = function () {}; // mengecek status
// xhr.onerror = function () {}; // menampilkan error ketika request
// xhr.send();
// xhr -> kepanjangan XMLHttpRequest

const promise = new Promise((resolve, reject) => {
  const xhttp = new XMLHttpRequest();

  xhttp.open("GET", "https://api.npoint.io/07c1d6487059fd250aa1", true);
  xhttp.onload = function () {
    if (xhttp.status === 200) {
      resolve(JSON.parse(xhttp.responseText));
    } else if (xhttp.status === 400) {
      reject("Error loading data");
    }
  };
  xhttp.onerror = function () {
    reject("Network Error");
  };
  xhttp.send();
});

// console.log(promise);

// asynch-await
let dataTestimonial = [];
async function getData(rating) {
  try {
    const response = await promise;
    console.log(response);
    dataTestimonial = response;
    testimonials();
  } catch (error) {
    console.log(error);
  }
}
getData();

function testimonials() {
  let testimonialHtml = "";

  dataTestimonial.forEach((card) => {
    testimonialHtml += `<div class="card-test" >
         <div class="card-title-test">
             <img src="${card.img}" alt="" />
         </div>
         <div class="card-body-test">
             <div class="card-isi">
                 <p>"${card.words}"</p>
             </div>
             <div class="card-pengisi">
               <p>- ${card.user}</p>
               <i class="fa-solid fa-star"></i>
               <p>${card.rating}</p>
             </div>
         </div>
     </div>`;
  });
  document.getElementById("testimonial").innerHTML = testimonialHtml;
}

// fitur rating
function testimonialFilter(rating) {
  let testimonialHtmlFilter = "";
  const dataFilter = dataTestimonial.filter((card) => {
    return card.rating === rating;
  });
  dataFilter.forEach((card) => {
    testimonialHtmlFilter += `<div class="card-test" >
       <div class="card-title-test">
           <img src="${card.img}" alt="" />
       </div>
       <div class="card-body-test">
           <div class="card-isi">
               <p>"${card.words}"</p>
           </div>
           <div class="card-pengisi">
             <p>- ${card.user}</p>
             <i class="fa-solid fa-star"></i>
             <p>${card.rating}</p>
           </div>
       </div>
   </div>`;
  });
  document.getElementById("testimonial").innerHTML = testimonialHtmlFilter;
}
