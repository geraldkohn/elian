import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { RootState, AppThunk } from "../../app/store";

export interface RegisterState {
  value: number;
  status: "idle" | "loading" | "failed";
}

const initialState: RegisterState = {
  value: 0,
  status: "idle",
};

export const registerSlice = createSlice({
  name: "register",
  initialState,
  reducers: {
    increment: (state) => {
      state.value += 1;
    },
    decrement: (state) => {
      state.value -= 1;
    },
    incrementByAmount: (state, action: PayloadAction<number>) => {
      state.value += action.payload;
    },
  },
});

export const { increment, decrement, incrementByAmount } =
  registerSlice.actions;

export const selectCount = (state: RootState) => state.register.value;
export const incrementIfOdd =
  (amount: number): AppThunk =>
  (dispatch, getState) => {
    const currentValue = selectCount(getState());
    if (currentValue % 2 === 1) {
      dispatch(incrementByAmount(amount));
    }
  };

export default registerSlice.reducer;
