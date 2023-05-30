// Add Blog

let dataBlog = [];

function addBlog(event) {
  event.preventDefault();

  let title = document.getElementById("title").value;
  let start = new Date(document.getElementById("start").value);
  let end = new Date(document.getElementById("end").value);
  let desc = document.getElementById("desc").value;
  let image = document.getElementById("upload").files;

  

  const jsIcon = '<i class="fa-brands fa-square-js fa-xl fa-fw"></i>';
  const reactIcon = '<i class="fa-brands fa-react fa-xl fa-fw"></i>';
  const phpIcon = '<i class="fa-brands fa-php fa-xl fa-fw"></i>';
  const javaIcon = '<i class="fa-brands fa-java fa-xl fa-fw"></i>';

  let diff = new Date(end - start);
  let mounth = diff.getMonth();
  let days = diff.getDate();
  

  let js = document.getElementById("tech_js").checked ? jsIcon : "";
  let react = document.getElementById("tech_react").checked ? reactIcon : "";
  let php = document.getElementById("tech_php").checked ? phpIcon : "";
  let java = document.getElementById("tech_java").checked ? javaIcon : "";

  image = URL.createObjectURL(image[0]);
  console.log(image);

  let blog = {
    title,
    mounth,
    days,
    desc,
    js,
    react,
    php,
    java,
    image,
    postAt: new Date(),
    author: "fernando",
  };

  dataBlog.push(blog);
  console.log(dataBlog);
  renderBlog();
}

// Time
function getFullTime(time) {
  let monthName = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];

  let date = time.getDate();

  let monthIndex = time.getMonth();

  let year = time.getFullYear();
  let hours = time.getHours();
  let minutes = time.getMinutes();
  // console.log(minutes);

  if (hours <= 9) {
    hours = "0" + hours;
  } else if (minutes <= 9) {
    minutes = "0" + minutes;
  }

  return `${date} ${monthName[monthIndex]} ${year} ${hours}:${minutes} WIB`;
}




function getDistanceTime(time) {
  let timeNow = new Date();
  let timePost = time;

  // waktu sekarang - waktu post
  let distance = timeNow - timePost; // hasilnya milidetik
  console.log(distance);

  let milisecond = 1000; // milisecond
  let secondInHours = 3600; // 1 jam 3600 detik
  let hoursInDays = 24; // 1 hari 24 jam

  let distanceDay = Math.floor(distance / (milisecond * secondInHours * hoursInDays)); // 1/86400000
  let distanceHours = Math.floor(distance / (milisecond * 60 * 60)); // 1/3600000
  let distanceMinutes = Math.floor(distance / (milisecond * 60)); // 1/60000
  let distanceSeconds = Math.floor(distance / milisecond); // 1/1000

  if (distanceDay > 0) {
    return `${distanceDay} Day Ago`;
  } else if (distanceHours > 0) {
    return `${distanceHours} Hours Ago`;
  } else if (distanceMinutes > 0) {
    return `${distanceMinutes} Minutes Ago`;
  } else {
    return `${distanceSeconds} Seconds Ago`;
  }
}

setInterval(function () {
  renderBlog();
}, 3000);

function renderBlog() {
  document.getElementById("content").innerHTML = "";

  x = 0;
  for (let i = x; i < dataBlog.length; i++) {
    document.getElementById("content").innerHTML += `
    <div class="project__card">
          <div class="project__card__img">
            <img src="${dataBlog[i].image}" alt="#" />
          </div>
          <div class="project__card__body">
            <a href="/content/project.html" class="title__project"><h4>${dataBlog[i].title}</h4></a>
            <p class="durasi">Durasi: ${dataBlog[i].mounth} bulan, ${dataBlog[i].days} hari</p>
            <p class="body">${dataBlog[i].desc}</p>
            <div class="flex iconTech">
            ${dataBlog[i].js}
            ${dataBlog[i].react}
            ${dataBlog[i].php}
            ${dataBlog[i].java}
            </div>
            <div class="btn__card">
              <button class="edit">Edit</button>
              <button class="delete">Delete</button>
            </div>
            <div style="float: right; margin: 10px">
                <p style="font-size: 15px; color: grey">${getDistanceTime(dataBlog[i].postAt)}</p>
              </div>
          </div>
        </div>
    `;
  }
}
