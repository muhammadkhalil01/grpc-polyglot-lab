const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const path = require("path");

const PROTO_PATH = path.join(__dirname, "../user.proto");

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

const userProto = grpc.loadPackageDefinition(packageDefinition).user;

function main() {
  const client = new userProto.UserService(
    "localhost:50051",
    grpc.credentials.createInsecure()
  );

  console.log("Mengirim request ke Go Server menggunakan Node.js...");
  client.GetUserProfile({ user_id: 1 }, (error, response) => {
    if (error) {
      console.error("Terjadi error:", error.message);
      client.close();
      process.exitCode = 1;
      return;
    }

    console.log(
      `Respons dari Server -> Nama: ${response.name}, ` +
      `Email: ${response.email}, Telepon: ${response.phone}`
    );
    client.close();
  });
}

main();
