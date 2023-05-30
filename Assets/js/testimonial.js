// // Superclass
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
//     return `<div class="testimonial">
//         <img
//             src="${this.image}"
//             class="profile-testimonial"
//         />
//         <p class="quote">${this.quote}</p>
//         <p class="author">- ${this.author}</p>
//     </div>
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

// let testiData = [testi1, testi2, testi3];
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
    author: "Rai Rai",
    quote: "keren",
    image: "https://pbs.twimg.com/media/FseC4rRaYAEIrU8?format=jpg&name=large",
    rating: 3,
  },
  {
    author: "Rai Rai",
    quote: "keren",
    image: "https://pbs.twimg.com/media/FseC4rRaYAEIrU8?format=jpg&name=large",
    rating: 4,
  },
  {
    author: "Rai Rai",
    quote: "keren",
    image: "https://pbs.twimg.com/media/FseC4rRaYAEIrU8?format=jpg&name=large",
    rating: 3,
  },
  {
    author: "Rai Rai",
    quote: "keren",
    image: "https://pbs.twimg.com/media/FseC4rRaYAEIrU8?format=jpg&name=large",
    rating: 5,
  },
];

function allTestimonials() {
  let testimonialHTML = "";

  testiData.forEach(function (item) {
    testimonialHTML += `<div class="testimonial">
            <img
                src="${item.image}"
                class="profile-testimonial"
            />
            <p class="quote">${item.quote}</p>
            <p class="author">- ${item.author}</p>
            <p class="rate">${item.rating} <i class="fa-solid fa-star"></i></p>
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
    testimonialHTML += `<h1>Data Not Found</h1>`;
  } else {
    testimonialFiltered.forEach(function (item) {
      testimonialHTML += `<div class="testimonial">
      <img
          src="${item.image}"
          class="profile-testimonial"
      />
      <p class="quote">${item.quote}</p>
      <p class="author">- ${item.author}</p>
      <p class="rate">${item.rating} <i class="fa-solid fa-star"></i></p>
  </div>
`;
    });
  }

  document.getElementById("testimonials").innerHTML = testimonialHTML;
}
