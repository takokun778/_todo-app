import * as User from './api/user/index';

(async function() {
    const basePath = 'http://localhost:8080';

    const conf = new User.Configuration({ basePath });

    const userApi = new User.UserServiceApi(conf);

    console.log('register');
    const registerRequest: User.V1RegisterRequest = {
        firstName: '玲',
        lastName: '大園',
        birthday: {
            year: 2001,
            month: 4,
            day: 18,
        },
    };

    const registerResponse = await userApi.userServiceRegister(registerRequest);

    console.log(registerResponse.data);

    console.log('get');
    const getResponse = await userApi.userServiceGet(registerResponse.data.user?.id!);

    console.log(getResponse.data);

    console.log('update');
    const updateRequest: User.InlineObject = {
        firstName: 'れい',
        lastName: 'おおぞの',
    };

    const updateResponse = await userApi.userServiceUpdate(getResponse.data.user?.id!, updateRequest);

    console.log(updateResponse.data);

    console.log('terminate');

    const deleteResponse = await userApi.userServiceTerminate(updateResponse.data.user?.id!);

    console.log(deleteResponse.data);
})();
