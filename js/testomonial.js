// let me = {
//   name: "fajar rohino",
//   profesi: "fullstack development",
// };
// console.log(me);

// class, perents
// class dataTestimonial {
//   #words = "";
//   #img = "";
//   star = "";

//   constructor(words, img, star) {
//     this.#words = words;
//     this.#img = img;
//     this.star = star;
//   }
//   get words() {
//     return this.#words;
//   }
//   get img() {
//     return this.#img;
//   }
//   get user() {
//     throw new Error("there is must be user to make testimonials!");
//   }
//   get star() {
//     return this.star;
//   }
// }
// // child
// class userTestimonial extends dataTestimonial {
//   #user = "";
//   constructor(user, words, img, star) {
//     super(words, img, star);
//     this.#user = user;
//   }
//   get user() {
//     return this.#user + " : User";
//   }
// }
// // child
// class companyTestimonial extends dataTestimonial {
//   #company = "";
//   constructor(company, words, img, star) {
//     super(words, img, star);
//     this.#company = company;
//   }
//   get user() {
//     return this.#company + " : Company";
//   }
// }
// const testim1 = new userTestimonial("ino", "keren banger", "https://images.unsplash.com/photo-1633332755192-727a05c4013d?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=580&q=80", "4");
// //
// const testim2 = new userTestimonial("inoy", "alhamdulillah", "https://images.unsplash.com/photo-1591258739299-5b65d5cbb235?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80", "5");
// //
// const testim3 = new userTestimonial("inoy", "alhamdulillah", "https://images.unsplash.com/photo-1591258739299-5b65d5cbb235?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80", "5");
// //
// const testim4 = new userTestimonial("inoy", "alhamdulillah", "https://images.unsplash.com/photo-1591258739299-5b65d5cbb235?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80", "4");
// //
// const testim5 = new companyTestimonial("eannoy", "bolleh bolleh", "https://images.unsplash.com/photo-1611485988300-b7530defb8e2?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80", "2");
// //
// const testim6 = new companyTestimonial("eannoy", "bolleh bolleh", "https://images.unsplash.com/photo-1611485988300-b7530defb8e2?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80", "1");
// //
// const testim7 = new companyTestimonial("eannoy", "bolleh bolleh", "https://images.unsplash.com/photo-1611485988300-b7530defb8e2?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80", "1");

// let arryTestimonialData = [testim1, testim2, testim3, testim4, testim5, testim6, testim7];
// let htmlTestimonial = "";
// for (let i = 0; i < arryTestimonialData.length; i++) {
//   htmlTestimonial += `
//     <div class="card-test" >
//         <div class="card-title-test">
//             <img src="${arryTestimonialData[i].img}" width="300px" alt="" />
//         </div>
//         <div class="card-body-test">
//             <div class="card-isi">
//                 <p>"${arryTestimonialData[i].words}"</p>
//             </div>c
//             <div class="card-pengisi">
//               <p>- ${arryTestimonialData[i].user}</p>
//               <i class="fa-solid fa-star"></i>
//               <p>${arryTestimonialData[i].star}</p>
//             </div>
//         </div>
//     </div>`;
// }
// document.getElementById("testimonial").innerHTML = htmlTestimonial;

// semua data di simpan di sini dalam bentuk array of object
const dataTestimonial = [
  {
    user: "Fajar Rohino",
    words: "Wow Amazing!",
    img: "https://images.unsplash.com/photo-1633332755192-727a05c4013d?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=580&q=80",
    rating: 5,
  },
  {
    user: "Rohino",
    words: "Hmmm Lumayan",
    img: "https://images.unsplash.com/photo-1591258739299-5b65d5cbb235?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80",
    rating: 3,
  },
  {
    user: "Ino",
    words: "This Website Cool!",
    img: "https://images.unsplash.com/photo-1503919545889-aef636e10ad4?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80",
    rating: 4,
  },
  {
    user: "eannoy",
    words: "Apaan nih",
    img: "https://images.unsplash.com/photo-1611485988300-b7530defb8e2?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80",
    rating: 2,
  },
  {
    user: "Fajar",
    words: "Hmmm....",
    img: "https://images.unsplash.com/photo-1577975882846-431adc8c2009?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=580&q=80",
    rating: 1,
  },
  {
    user: "noname",
    words: "Kurang bagus!",
    img: "https://images.unsplash.com/photo-1488371934083-edb7857977df?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=380&q=80",
    rating: 1,
  },
  {
    user: "Inox",
    words: "Lumayan lah",
    img: "https://images.unsplash.com/photo-1633332755192-727a05c4013d?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=580&q=80",
    rating: 3,
  },
];
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
// utk menampilkan semua atau secara default
testimonials();

// utk menampilkan rating get by ID di HTML id:testimonialFilter(prameter)
function testimonialFilter(rating) {
  let testimonialHtmlFilter = "";

  // menggunakan bulid.HOF filter
  const dataFilter = dataTestimonial.filter((card) => {
    return card.rating === rating;
  });
  // menggunakan bulid.HOF forEach karena looping dan tidak me-return data
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
