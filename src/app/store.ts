// 引入configureStore创建全局store
import { configureStore} from '@reduxjs/toolkit';
import { ThunkAction, Action } from '@reduxjs/toolkit';
// 从自定一的Slice引入自定义的Reducer
import loginReducer from '../features/login/loginSlice';
import registerReducer from '../features/register/registerSlice'
import homeReducer from '../features/home/homeSlice';

export const store = configureStore({
  reducer: {
    login: loginReducer,
    register: registerReducer,
    home: homeReducer
  },
});


export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
