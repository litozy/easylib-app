const Home = async () => {
  const response = await fetch(`${process.env.API_BASE_URL}/member`);
  const member = await response.json();
  console.log(member);

  return (
    <div>
      <h1>Testing mantap</h1>
    </div>
  );
};

export default Home;
