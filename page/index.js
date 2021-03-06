function Count() {
   const text = document.getElementById('textarea').value;

   const data = {'text': text};

   fetch('/count', {
      method: 'POST',
      headers: {
         'Content-Type': 'application/json'
      },
      cors: 'no-cors',
      body: JSON.stringify(data),
   })
   .then(res => res.json())
   .then(function(resJSON) {
      console.log(resJSON['Object'])
      document.getElementById('result_withspace').value = resJSON['withspace']
      document.getElementById('result_nospace').value = resJSON['nospace']
   });
}
