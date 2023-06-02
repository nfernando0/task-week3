// Superclass
// class Testimonial {
//   #quote = "";
//   #image = "";

//   constructor(quote, image) {
//     this.#quote = quote;
//     this.#image = image;
//   }

//   get quote() {
//     return this.#quote;
//   }

//   get image() {
//     return this.#image;
//   }

//   get author() {
//     throw new Error("getAuthor() method must be implemented");
//   }

//   get testimonialHTML() {
//     return `<div class="col-lg-4">
//     <div class="card my-4 mx-auto" style="width: 20rem">
//       <img src="${this.image}" class="card-img-top" />
//       <div class="card-body">
//         <p>${this.quote}</p>
//         <p class="float-end">- ${this.author} -</p>
//       </div>
//     </div>
//   </div>
// `;
//   }
// }

// // SubClass

// class AuthorTestimonials extends Testimonial {
//   #author = "";

//   constructor(author, quote, image) {
//     super(quote, image);
//     this.#author = author;
//   }

//   get author() {
//     return this.#author;
//   }
// }

// class CompanyTestimonials extends Testimonial {
//   #company = "";

//   constructor(company, quote, image) {
//     super(quote, image);
//     this.#company = company;
//   }

//   get author() {
//     return this.#company + " Company";
//   }
// }

// const testi1 = new AuthorTestimonials("Rai rai", "keren cuyyy", "https://pbs.twimg.com/media/FseC4rRaYAEIrU8?format=jpg&name=large");
// const testi2 = new AuthorTestimonials("Rai rai", "keren cuyyy", "https://pbs.twimg.com/media/FseC4rRaYAEIrU8?format=jpg&name=large");
// const testi3 = new AuthorTestimonials("Rai rai", "keren cuyyy", "https://pbs.twimg.com/media/FseC4rRaYAEIrU8?format=jpg&name=large");
// const testi4 = new AuthorTestimonials("Rai rai", "keren cuyyy", "https://pbs.twimg.com/media/FseC4rRaYAEIrU8?format=jpg&name=large");

// let testiData = [testi1, testi2, testi3, testi4];
// let testimonialHTML = "";

// for (let i = 0; i < testiData.length; i++) {
//   testimonialHTML += testiData[i].testimonialHTML;
// }

// document.getElementById("testimonials").innerHTML = testimonialHTML;

const testiData = [
  {
    author: "Rai Rai",
    quote: "keren",
    image: "https://pbs.twimg.com/media/FseC4rRaYAEIrU8?format=jpg&name=large",
    rating: 5,
  },
  {
    author: "azusa",
    quote: "keren banget",
    image: "https://pbs.twimg.com/media/FudBlYVaQAARz0s?format=jpg&name=large",
    rating: 5,
  },
  {
    author: "Rai Rai",
    quote: "keren",
    image: "https://static.vecteezy.com/system/resources/previews/003/586/230/original/no-photo-sign-sticker-with-text-inscription-on-isolated-background-free-vector.jpg",
    rating: 4,
  },
  {
    author: "Rai Rai",
    quote: "keren",
    image: "https://pbs.twimg.com/media/FuzxrTjagAARYJs?format=jpg&name=large",
    rating: 3,
  },
  {
    author: "koharu",
    quote: "keren",
    image: "https://pbs.twimg.com/media/Fu2gTlqakAARjp9?format=jpg&name=small",
    rating: 5,
  },
];

function allTestimonials() {
  let testimonialHTML = "";

  testiData.forEach(function (item) {
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
allTestimonials();

function filterTestimonials(rating) {
  let testimonialHTML = "";

  const testimonialFiltered = testiData.filter(function (item) {
    return item.rating === rating;
  });

  if (testimonialFiltered.length === 0) {
    testimonialHTML += `<h1 class="text-center my-3">Data Not Found</h1>`;
  } else {
    testimonialFiltered.forEach(function (item) {
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
