/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hirokiurabe <hirokiurabe@student.42.fr>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2025/02/21 11:18:28 by hirokiurabe       #+#    #+#             */
/*   Updated: 2025/02/21 11:45:30 by hirokiurabe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
)

func main() {
	// args check
	if len(os.Args) != 2 {
		fmt.Println("error: invalid argument")
		return
	}
	fmt.Println("Hello World!")
}
