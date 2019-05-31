# roll-the-dices

A simple personal project to catch up with Golang.
Currently only a bunch of scripts that make some queries to mySQL, some requests to an API and calculates some random results from dices.

## Features
Current accepted commands and parameters, returned from API requests:

### Parameters
**dice**: [`operator`] `dice_number` d `die faces`
>Both _dice number_ and _die faces_ should be integers.  _Operator_ (*+*) should be indicated if there is more than one dice in the command.
>
_Examples_: 1d20, 1d200, 1d8+3d6

**bonus**: `operator` `bonus quantifier`
>_Operator_ (*+* or *-*) should be indicated always, even if the command didn't use _dice_ parameters.
>
_Examples_: +7. +3, -10, +1+3-2

**tag**: `text`
>_Text_ could include any text but special symbols or _dice_ or _bonus_ nomenclature.
>
_Examples_: Initiative, test, searching for someone, etc

### Commands
**tira** [_dice_[_dice_[...]]] [_bonus_[_bonus_[...]]] [_tag_]
> Resolves _**dice**_ parameters, if any. Default value is "1d20". Then sum _**bonus**_ parameters to the result, if any. If _**tag**_ parameters are passed they will be included in the beginning of the roll, working as a title or tag.
> 
_Examples:_
| command | result
| --- | ---
| tira 1d10+7 | 1d10[7]+7= 14
| tira +2 | 1d20[12]+2 = 14
| tira 1d20+3 Initiative | Initiative: 1d20[2]+3= 5

**v** [_bonus_[_bonus_[...]]] [_tag_]
> Resolves an advantage roll: Rolling two twenty-sided dice and taking the higher roll
> 
_Examples:_
| command | result
| --- | ---
| v +7 | 2d10[7 12]+7= 19
| v -1 | 2d20[12 19]-1 = 18
| v Initiative | Initiative: 1d20[2 7]= 7

**dv** [_bonus_[_bonus_[...]]] [_tag_]
> Resolves a disadvantage roll: Rolling two twenty-sided dice and taking the lower roll
> 
_Examples:_
| command | result
| --- | ---
| dv +7 | 2d10[7 12]+7= 14
| dv -1 | 2d20[12 19]-1 = 11
| dv Initiative | Initiative: 1d20[2 7]= 2

## Versioning
Current stable version: 0.2.2-beta

## License
This project is licensed under GPL-3.0, see [LICENSE](./LICENSE) file for details
