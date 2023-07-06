function submitData(event) {
  event.preventDefault();

  let name = document.getElementById("input-name").value;
  let email = document.getElementById("input-email").value;
  let phone = document.getElementById("input-phone").value;
  let subject = document.getElementById("input-subject").value;
  let message = document.getElementById("input-message").value;
  // javascript Regex untuk variabel untuk nomor
  var number = /^[0-9]+$/;
  // javascript Regex untuk variabel untuk email
  var mailformat = /^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$/;

  //   console.log(name);
  //   console.log(email);
  //   console.log(phone);
  //   console.log(subject);
  //   console.log(message);

  //   alert(
  //     `name : ${name} \n
  //       email : ${email} \n
  //       phone : ${phone} \n
  //       subject : ${subject} \n
  //       message : ${message} \n`
  //   );

  let objectData = {
    name,
    email,
    phone,
    subject,
    message,
  };
  let arrayData = [name, email, phone, subject, message];
  console.log(objectData);
  // validation form
  if (name == "" && email == "" && phone == "" && subject == "" && message == "") {
    return alert("Form harus diisi!");
  } else if (name === "") {
    return alert("Nama harus diisi!");
  } else if (email === "") {
    return alert("Email harus diisi!");
  } else if (!email.match(mailformat)) {
    return alert("Email harus sesuai dengan format!");
  } else if (phone == "") {
    return alert("Phone number harus diisi!");
  } else if (!phone.match(number)) {
    return alert("Phone number harus diisi dengan angka!");
  } else if (phone.length != 12) {
    return alert("Phone number harus 12 digit!");
  } else if (subject === "") {
    return alert("Subject harus diisi!");
  } else if (message === "") {
    return alert("Pesan harus diisi!");
  }
  const emailReceiver = "fajarshadow20@gmail.com";
  let a = document.createElement("a");
  a.href = `mailto:${emailReceiver}?subject=${subject}&body= Hy my name is ${name}, \n ${message} \n please contact me ${phone}`;
  a.click();
}
