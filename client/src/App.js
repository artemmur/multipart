import './App.css';

const list = []

const send = (e) => {
  e.preventDefault();
  const form = new FormData();
  list.forEach(async (item) => {
    form.set('file', item);
    await fetch('http://localhost:8888/', {
      body: form,
      method: 'POST',
      mode: 'no-cors',
    });
  });
}

function App() {
  return (
    <form onSubmit={(e) => { send(e) }}>
      <input id='file-input' type='file' multiple onChange={(e) => {
        e.preventDefault();
        for (var i = 0; i < e.target.files.length; i++) {
          list.push(e.target.files[i]);
        }
      }} />
      <button type='submit'>
        Submit
      </button>
    </form>
  );
}

export default App;
