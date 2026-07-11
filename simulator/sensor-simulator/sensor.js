export function generateReading() {

    return {

        deviceId: "TEMP001",

        temperature: +(25 + Math.random() * 10).toFixed(2),

        humidity: Math.floor(40 + Math.random() * 40),

        battery: Math.floor(60 + Math.random() * 40)

    };

}