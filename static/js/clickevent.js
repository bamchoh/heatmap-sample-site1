let body = document.getElementById("body");

body.addEventListener('click', function(e) {
  data = {
    url: location.href,
    x: e.pageX / document.documentElement.scrollWidth,
    y: e.pageY / document.documentElement.scrollHeight
  };

  fetch('https://radiant-spire-12330.herokuapp.com/api/click', {
    method: "POST",
    mode: "no-cors",
    cache: "no-cache",
    credentials: "same-origin",
    headers: {
      "Content-Type": "application/json; charset=utf-8",
    },
    redirect: "follow",
    referrer: "no-referrer",
    body: JSON.stringify(data),
  });
});
