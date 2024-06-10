document.getElementById('submitBtn').addEventListener('click', () => {
  const form = document.getElementById('coeficientsForm');
  const formData = new FormData(form);
  const data = {};

  formData.forEach((value, key) => {
      data[key] = parseInt(value, 10);
  });

  fetch("http://localhost:8080/", {
      method: "POST",
      headers: {
          'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
  })
  .then(response => {
      if (!response.ok) {
          throw new Error('Network response was not ok ' + response.statusText);
      }
      return response.json();
  })
  .then(data => {
      console.log('Success:', data);
  })
  .catch((error) => {
      console.error('Error:', error);
  });
});
