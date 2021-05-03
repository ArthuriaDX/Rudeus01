// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package pTools

// Flag is the options passed along with the commands
// by users. they should send them with prefex "--",
// but we will remove them in the pTools.
type Flag string

// Arg is the arguments for commands
type Arg []string
