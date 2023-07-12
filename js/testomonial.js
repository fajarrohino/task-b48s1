// let me = {
//   name: "fajar rohino",
//   profesi: "fullstack development",
// };
// console.log(me);

// class, perents
class dataTestimonial {
  #words = "";
  #img = "";

  constructor(words, img) {
    this.#words = words;
    this.#img = img;
  }
  get words() {
    return this.#words;
  }
  get img() {
    return this.#img;
  }
  get user() {
    throw new Error("there is must be user to make testimonials!");
  }
}
// child
class userTestimonial extends dataTestimonial {
  #user = "";
  constructor(user, words, img) {
    super(words, img);
    this.#user = user;
  }
  get user() {
    return this.#user + " : User";
  }
}
// child
class companyTestimonial extends dataTestimonial {
  #company = "";
  constructor(company, words, img) {
    super(words, img);
    this.#company = company;
  }
  get user() {
    return this.#company + " : Company";
  }
}
const testim1 = new userTestimonial("ino", "keren banger", "https://images.unsplash.com/photo-1633332755192-727a05c4013d?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=580&q=80");
//
const testim2 = new userTestimonial("inoy", "alhamdulillah", "https://images.unsplash.com/photo-1591258739299-5b65d5cbb235?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80");
//
const testim3 = new companyTestimonial("eannoy", "bolleh bolleh", "https://images.unsplash.com/photo-1611485988300-b7530defb8e2?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=387&q=80");

let arryTestimonialData = [testim1, testim2, testim3];
let htmlTestimonial = "";
for (let i = 0; i < arryTestimonialData.length; i++) {
  htmlTestimonial += `
    <div class="card-test" >
        <div class="card-title-test">
            <img src="${arryTestimonialData[i].img}" width="300px" alt="" />
        </div>
        <div class="card-body-test">
            <div class="card-isi">
                <p>"${arryTestimonialData[i].words}"</p>
            </div>
            <div class="card-pengisi">
                <p>- ${arryTestimonialData[i].user}</p>
            </div>
        </div>
    </div>`;
}
document.getElementById("testimonial").innerHTML = htmlTestimonial;
