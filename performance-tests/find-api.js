// import necessary modules
import { check } from "k6";
import http from "k6/http";

// define configuration
export const options = {
  scenarios: {
    ramping_load: {
      executor: "ramping-vus",
      stages: [
        { duration: "30s", target: 10 },
        { duration: "30s", target: 1000 },
        { duration: "2m", target: 1000 }
      ]
    },
    // ramping_load: {
    //   executor: 'ramping-vus',
    //   startVUs: 1,
    //   stages: [
    //     { duration: '1m', target: 1000 },
    //     { duration: '1m', target: 1000 },
    //     { duration: '1m', target: 1 },
    //   ],
    //   gracefulRampDown: '5s'
    // }
  },
};

class Random {
  constructor(){
    this.SEQ = [1,30, 100, 300, 99, 50, 500, 2, 4, 5, 8 ,21, 15, 71, 82];
    this.index = 0;
  }

  next() {
    const value = this.SEQ[this.index];
    this.index = this.index % this.SEQ.length;
    return value;
  }
}

const random = new Random()

export default function () {
  const start = new Date(1985, 0, 1).getTime();
  const end = new Date(2000, 0, 28).getTime();
  const randomDate = new Date(start + random.next() * (end - start));
  const formattedDate = `${randomDate.getFullYear()}-${String(
    randomDate.getMonth() + 1
  ).padStart(2, "0")}-${String(randomDate.getDate()).padStart(2, "0")}`;
  const url = `http://localhost:${__ENV.PORT}/employees?after=${formattedDate}`;
  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  };
  const res = http.get(url, params);
  check(res, {
    "response code was 200": (res) => res.status == 200,
  });
}
