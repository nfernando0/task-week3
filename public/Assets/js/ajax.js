const promise = new Promise((resolve, reject) => {
  const xhr = new XMLHttpRequest();
  xhr.open("GET", "https://api.npoint.io/62b40e4f92a7d6563561", true);
  //   console.log(xhr);
  xhr.onload = () => {
    if (xhr.status === 200) {
      // We parsing it so it is easier to read in console
      // response vs responseText, the differences are, responseText is an older version, when response is more newer, but the output is still the same/similiar.
      resolve(JSON.parse(xhr.response));
    } else {
      reject("Error loading data.");
    }
  };
  xhr.onerror = () => {
    reject("Network error.");
  };
  xhr.send();
});

async function getAllTestimonials() {
  const response = await promise;
  //   console.log(response);

  let testimonialHTML = "";
  response.forEach(function (item) {
    testimonialHTML += `<div class="col-lg-4">
      <div class="card my-4 mx-auto" style="width: 20rem">
        <img src="${item.image}" class="card-img-top" />
        <div class="card-body">
          <p>${item.quote}</p>
          <p class="float-end">- ${item.author} -</p>
          <p class="rate">${item.rating} <i class="fa-solid fa-star"></i></p>
        </div>
      </div>
    </div>
  `;
  });

  document.getElementById("testimonials").innerHTML = testimonialHTML;
}

getAllTestimonials();

async function getFilteredTestimonials(rating) {
  const response = await promise;

  const testimonialFiltered = response.filter((item) => {
    return item.rating === rating;
  });

  //   console.log(testimonialFiltered);

  let testimonialHTML = "";

  if (testimonialFiltered.length === 0) {
    testimonialHTML = "<h1>Data not found!</h1>";
  } else {
    testimonialFiltered.forEach((item) => {
      testimonialHTML += `<div class="col-lg-4">
      <div class="card my-4 mx-auto" style="width: 20rem">
        <img src="${item.image}" class="card-img-top" />
        <div class="card-body">
          <p>${item.quote}</p>
          <p class="float-end">- ${item.author} -</p>
          <p class="rate">${item.rating} <i class="fa-solid fa-star"></i></p>
        </div>
      </div>
    </div>
  `;
    });
  }

  document.getElementById("testimonials").innerHTML = testimonialHTML;
}
