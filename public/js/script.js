const form = document.querySelector('form#upload');

form.addEventListener('submit', evt => {
  evt.preventDefault();

  const { files } = form.querySelector('input[type=file]');

  if (!files.length) {
    alert('Select an image');
  }

  const img = files[0];
  const formData = new FormData();
  formData.append('image', img);

  fetch('/api/upload', {
    method: 'post',
    body: formData
  });

  return false;
});
