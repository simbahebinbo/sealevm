/*
 * Copyright 2020 The SealEVM Authors
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package instructions

import (
	"SealEVM/common"
	"SealEVM/evmInt256"
	"SealEVM/opcodes"
)

func loadStack() {
	instructionTable[opcodes.POP] = opCodeInstruction {
		action: func(ctx *instructionsContext) (bytes []byte, err error) {
			_, _ = ctx.stack.Pop()
			return nil, nil
		},
		minStackDepth: 1,
		enabled:       true,
	}

	setPushActions()
	setSwapActions()
	setDupActions()
}

func setPushActions() {
	for i := opcodes.PUSH1; i <= opcodes.PUSH32; i++ {
		bytesSize := uint64(i - opcodes.PUSH1 + 1)

		instructionTable[i] = opCodeInstruction {
			action: func(ctx *instructionsContext) ([]byte, error) {
				start := ctx.pc + 1

				codeBytes := common.GetDataFrom(ctx.environment.Contract.Code, start, bytesSize)

				i := evmInt256.New(0)
				i.SetBytes(codeBytes)
				err := ctx.stack.Push(i)

				ctx.pc += bytesSize
				return nil, err
			},

			minStackDepth: 0,
			enabled:       true,
		}
	}
}

func setSwapActions()  {
	for i := opcodes.SWAP1; i <= opcodes.SWAP16; i ++ {
		swapDepth := int(i - opcodes.SWAP1 + 1)

		instructionTable[i] = opCodeInstruction {
			action: func(ctx *instructionsContext) ([]byte, error) {
				_ = ctx.stack.Swap(swapDepth)
				return nil, nil
			},

			minStackDepth: swapDepth + 1,
			enabled:       true,
		}
	}
}

func setDupActions()  {
	for i := opcodes.DUP1; i <= opcodes.DUP16; i ++ {
		dupDepth := int(i - opcodes.DUP1 + 1)

		instructionTable[i] = opCodeInstruction {
			action: func(ctx *instructionsContext) ([]byte, error) {
				err := ctx.stack.Dup(dupDepth)
				return nil, err
			},

			minStackDepth: dupDepth,
			enabled:       true,
		}
	}
}
